package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sainath123112/e-commerce-backend/user-service/model"
	"github.com/sainath123112/e-commerce-backend/user-service/pkg/jwt"
	"github.com/sainath123112/e-commerce-backend/user-service/service"
)

func Authenticate(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		vars := mux.Vars(r)
		userid, _ := uuid.Parse(vars["userid"])
		w.Header().Set("Content-Type", "application/json")
		if authorization == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(&model.ErrorResponse{Message: "No header authorization found", ErrorString: "Unauthorized"})
			return
		}
		token := authorization[len("Bearer "):]
		isValidToken, err := jwt.ValidateToken(token)
		if !isValidToken {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(&model.ErrorResponse{Message: "Invalid Token", ErrorString: err.Error()})
			return
		}
		emailFromToken, err := jwt.GetUsername(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(&model.ErrorResponse{Message: emailFromToken, ErrorString: err.Error()})
			return
		}
		emailFromUserId, err := service.GetUserEmail(userid)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(&model.ErrorResponse{Message: "No user found for id: " + vars["userid"], ErrorString: err.Error()})
			return
		}
		if emailFromUserId != emailFromToken {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(&model.ErrorResponse{Message: "token is not belongs to this user with user id: " + vars["userid"], ErrorString: "Token mismatch"})
			return
		}
		handler.ServeHTTP(w, r)
	}
}
