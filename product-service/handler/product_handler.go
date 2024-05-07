package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sainath123112/e-commerce-backend/product-service/model"
	"github.com/sainath123112/e-commerce-backend/product-service/service"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	queryParameters := r.URL.Query()["primary_category"]
	products, err := service.GetAllProductDetails(queryParameters)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&model.ErrorResponse{Message: "Unable to get data for products", ErrorString: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&products)
}

func GetSingleProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	product, err := service.GetProductWithId(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&model.ErrorResponse{Message: "Unable to get data for product with id: " + id, ErrorString: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&product)
}
