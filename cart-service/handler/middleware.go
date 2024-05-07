package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sainath123112/e-commerce-backend/cart-service/model"
	pb "github.com/sainath123112/e-commerce-backend/cart-service/user/proto"
	"google.golang.org/grpc"
)

func AuthHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalln("Unable to connect User Service due to: " + err.Error())
		}
		defer conn.Close()
		client := pb.NewUserServiceClient(conn)

		authorization := r.Header.Get("Authorization")
		vars := mux.Vars(r)
		userId := vars["id"]
		w.Header().Set("Content-Type", "application/json")
		if authorization == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(&model.ErrorResponse{Message: "No header authorization found", ErrorString: "Unauthorized"})
			return
		}
		authToken := authorization[len("Bearer "):]
		isValidToken, err := client.ValidateToken(context.TODO(), &pb.ValidateTokenRequest{Token: authToken})
		if err != nil && isValidToken == nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(&model.ErrorResponse{Message: "Invalid Token", ErrorString: err.Error()})
			return
		}

		emailFromToken, err := client.GetUserNameFromToken(context.TODO(), &pb.UserNameFromTokenRequest{Token: authToken})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(&model.ErrorResponse{Message: emailFromToken.EmailFromToken, ErrorString: err.Error()})
			return
		}

		emailFromUserId, err := client.GetUserNameFromUserId(context.TODO(), &pb.UserNameFromUserIdRequest{UserId: userId})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(&model.ErrorResponse{Message: "No user found for id: " + userId, ErrorString: err.Error()})
			return
		}

		if emailFromUserId.EmailFromUserId != emailFromToken.EmailFromToken {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(&model.ErrorResponse{Message: "token is not belongs to this user with user id: " + userId, ErrorString: "Token mismatch"})
			return
		}

		handler.ServeHTTP(w, r)
	}

}
