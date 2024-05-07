package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sainath123112/e-commerce-backend/cart-service/model"
	"github.com/sainath123112/e-commerce-backend/cart-service/service"
)

func GetCartItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId, err := uuid.Parse(vars["id"])
	if err != nil {
		log.Fatalln("Error parsing string to uuid due to: " + err.Error())
	}

	cartDetails, err := service.GetCartDetails(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "unable get cart details", "error": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(&cartDetails)
}

func AddCartItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cartItem model.CartItemRequest
	vars := mux.Vars(r)
	userId, _ := uuid.Parse(vars["id"])
	json.NewDecoder(r.Body).Decode(&cartItem)
	cart, err := service.GetCartDetails(userId)
	if err != nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(&model.ErrorResponse{Message: "Unable to get cart details", ErrorString: err.Error()})
		return
	}
	cartItems, err := service.AddItemToCart(int(cart.ID), cartItem)
	if err != nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(&model.ErrorResponse{Message: "Unable to add cart item", ErrorString: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(&cartItems)
}
