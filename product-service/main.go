package main

import (
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sainath123112/e-commerce-backend/product-service/config"
	"github.com/sainath123112/e-commerce-backend/product-service/handler"
	"github.com/sainath123112/e-commerce-backend/product-service/model/grpc_impl"
	pb "github.com/sainath123112/e-commerce-backend/product-service/proto"
	"google.golang.org/grpc"
)

func main() {
	//Creating TCP server
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalln("failed to listen due to: " + err.Error())
	}

	grpcServer := grpc.NewServer()

	pb.RegisterProductServiceServer(grpcServer, &grpc_impl.ProductServiceServer{})

	go func() {
		log.Println("Serving grpc on port: 50052")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalln("Unable to connect grpc server to tcp server")
		}
	}()

	ConfiYamlPath := "config/config.yaml"
	config.ReadConfigFile(ConfiYamlPath)
	Port := strconv.Itoa(config.ConfigObj.ProductService.Server.Port)
	pr := mux.NewRouter()
	product_routes := pr.PathPrefix("/v1/").Subrouter()
	product_routes.HandleFunc("/products", handler.GetAllProducts).Methods("GET")
	product_routes.HandleFunc("/products/{id}", handler.GetSingleProduct).Methods("GET")

	log.Println("Listening on port: " + Port)
	if err := http.ListenAndServe(":"+Port, pr); err != nil {
		log.Fatalln("Unable to lister at port: 8081 due to error: " + err.Error())
	}
}
