package model

type Product struct {
	Id                      string   `bson:"_id" json:"id"`
	Title                   string   `bson:"title" json:"title"`
	Brand                   string   `bson:"brand" json:"brand"`
	MainImage               string   `bson:"main_image" json:"main_image"`
	Description             string   `bson:"description" json:"description"`
	Currency                string   `bson:"currency" json:"currency"`
	Price                   float64  `bson:"price" json:"price"`
	Availability            string   `bson:"availability" json:"availability"`
	AvailableDeliveryMethod string   `bson:"availableDeliveryMethod" json:"available_delivery_method"`
	PrimaryCategory         string   `bson:"primary_category" json:"primary_category"`
	SubCategory1            string   `bson:"sub_category_1" json:"sub_category_1"`
	SubCategory2            string   `bson:"sub_category_2" json:"sub_category_2"`
	SubCategory3            string   `bson:"sub_category_3" json:"sub_category_3"`
	Images                  []string `bson:"images" json:"images"`
}

type ProductsResponseDto struct {
	Id              string  `bson:"_id" json:"id"`
	Title           string  `bson:"title" json:"title"`
	Brand           string  `bson:"brand" json:"brand"`
	MainImage       string  `bson:"main_image" json:"main_image"`
	Currency        string  `bson:"currency" json:"currency"`
	Price           float64 `bson:"price" json:"price"`
	Availability    string  `bson:"availability" json:"availability"`
	PrimaryCategory string  `bson:"primary_category" json:"primary_category"`
}
