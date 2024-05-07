package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sainath123112/e-commerce-backend/cart-service/handler"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/cart/{id}", handler.AuthHandler(handler.GetCartItems))
	r.HandleFunc("/cart/{id}/items", handler.AuthHandler(handler.AddCartItem)).Methods("POST")
	log.Println("Listening on port: 8083")
	http.ListenAndServe(":8083", r)
}
