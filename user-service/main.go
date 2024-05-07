package main

import (
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sainath123112/e-commerce-backend/user-service/handler"
	"github.com/sainath123112/e-commerce-backend/user-service/model/grpc_impl"
	pb "github.com/sainath123112/e-commerce-backend/user-service/proto"
	"google.golang.org/grpc"
)

func main() {

	//creating tcp server for gRPC service
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln("failed to listen due to: " + err.Error())
	}

	//creating gRPC server
	grpcServer := grpc.NewServer()
	//Register grpc server and server struct
	pb.RegisterUserServiceServer(grpcServer, &grpc_impl.UserServiceServer{})
	go func() {
		log.Println("Serving on port: 50051")
		//Serving grpcs server on tcp server port
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalln("Unable to listen on port 50051 due to: " + err.Error())
		}
	}()

	// go func() {
	// 	message := "Hello for "
	// 	kafka.UserServiceProducer(message)
	// }()

	//Creating Gorilla mux router to expose api end points
	r := mux.NewRouter()

	r.HandleFunc("/user/login", handler.LoginHandler).Methods("POST")
	r.HandleFunc("/user/register", handler.RegisterUser).Methods("POST")
	r.HandleFunc("/user/{userid}", handler.Authenticate(handler.GetUserDetails)).Methods("GET")
	http.Handle("/", r)
	log.Println("Listening and serving on port: 8082")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatalln(err.Error())
	}
}
