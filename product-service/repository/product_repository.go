package repository

import (
	"context"
	"log"

	"github.com/sainath123112/e-commerce-backend/product-service/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDbCollection() *mongo.Collection {
	ConfiYamlPath := "config/config.yaml"
	config.ReadConfigFile(ConfiYamlPath)
	uri := config.ConfigObj.ProductService.Database.Uri
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalln("Unable to connect Database due to error: " + err.Error())
	}
	collection := client.Database(config.ConfigObj.ProductService.Database.DatabaseName).Collection(config.ConfigObj.ProductService.Database.CollectionName)
	return collection
}
