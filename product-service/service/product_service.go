package service

import (
	"context"
	"errors"

	"github.com/sainath123112/e-commerce-backend/product-service/model"
	"github.com/sainath123112/e-commerce-backend/product-service/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collection *mongo.Collection
)

func init() {
	collection = repository.ConnectMongoDbCollection()
}

func GetAllProductDetails(queryParameters []string) ([]model.ProductsResponseDto, error) {
	var products []model.Product
	var productsResponseArray []model.ProductsResponseDto
	var productsResponse model.ProductsResponseDto
	var cursor *mongo.Cursor
	var err error
	filter := bson.D{{"primary_category", bson.D{{"$in", queryParameters}}}}
	if len(queryParameters) == 0 {
		cursor, err = collection.Find(context.TODO(), bson.D{})
	} else {
		cursor, err = collection.Find(context.TODO(), filter)
	}

	if err != nil {
		return []model.ProductsResponseDto{}, err
	}
	err = cursor.All(context.TODO(), &products)
	if err != nil {
		return []model.ProductsResponseDto{}, err
	}
	for _, product := range products {
		productsResponse.Id = product.Id
		productsResponse.Title = product.Title
		productsResponse.Brand = product.Brand
		productsResponse.MainImage = product.MainImage
		productsResponse.Currency = product.Currency
		productsResponse.Price = product.Price
		productsResponse.Availability = product.Availability
		productsResponse.PrimaryCategory = product.PrimaryCategory
		productsResponseArray = append(productsResponseArray, productsResponse)
	}
	return productsResponseArray, nil
}

func GetProductWithId(uniq_id string) (model.Product, error) {
	var product model.Product
	filter := bson.M{"_id": uniq_id}
	err := collection.FindOne(context.TODO(), filter).Decode(&product)
	if err != nil {
		return model.Product{}, err
	}
	return product, nil
}

func ValidateProductById(productId string) (bool, error) {
	var product model.Product
	filter := bson.M{"_id": productId}
	err := collection.FindOne(context.TODO(), filter).Decode(&product)
	if err != nil {
		return false, err
	}
	if product.Availability == "InStock" {
		return true, nil
	}
	return false, errors.New("product not in stock")
}
