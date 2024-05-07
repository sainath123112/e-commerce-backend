package service

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/sainath123112/e-commerce-backend/cart-service/model"
	productpb "github.com/sainath123112/e-commerce-backend/cart-service/product/proto"
	"github.com/sainath123112/e-commerce-backend/cart-service/repository"
	pb "github.com/sainath123112/e-commerce-backend/cart-service/user/proto"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	db, err = repository.ConnectDb()
	if err != nil {
		log.Fatalln("Unable to connect database due to: " + err.Error())
	}
	db.AutoMigrate(&model.Cart{}, &model.CartItem{})

}

func GetCartDetails(id uuid.UUID) (model.Cart, error) {
	var cartItems model.Cart
	idStr := id.String()
	isExist := GetIsUserExists(idStr)
	if isExist {
		err := db.Where("user_id = ?", id).Preload("CartItems").First(&cartItems).Error
		if err == gorm.ErrRecordNotFound {
			cartItems.UserId = id
			if err := db.Create(&cartItems).Error; err != nil {
				return model.Cart{}, err
			}
		}
		return cartItems, nil
	}
	return cartItems, errors.New("user not exist")

}

func GetIsUserExists(id string) bool {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Unable to connect User Service due to: " + err.Error())
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)

	res, err := client.GetUserExist(context.TODO(), &pb.GetUserRequest{Id: id})
	if err != nil {
		log.Fatalln("Unable to get response")
	}
	return res.IsExist
}

func AddItemToCart(cartId int, cartItemRequest model.CartItemRequest) (model.CartItem, error) {
	var cartItem model.CartItem
	isValid, err := ValidateProductById(cartItemRequest.ProductId.String())
	if !isValid {
		return model.CartItem{}, err
	}
	err = db.Model(&model.CartItem{}).Where("cart_id = ?", uint(cartId)).Where("product_id = ?", cartItemRequest.ProductId).First(&cartItem).Error
	if err == nil {
		err := db.Model(&model.CartItem{}).Where("cart_id = ?", uint(cartId)).Where("product_id = ?", cartItemRequest.ProductId).Update("quantity", cartItem.Quantity+cartItemRequest.Quantity).Error
		if err != nil {
			log.Fatalln("Unable to update cart item due to: " + err.Error())
		}
		db.Model(&model.CartItem{}).Where("cart_id = ?", uint(cartId)).Where("product_id = ?", cartItemRequest.ProductId).First(&cartItem)
		return cartItem, nil
	}
	cartItem.CartId = uint(cartId)
	cartItem.ProductId = cartItemRequest.ProductId
	cartItem.Quantity = cartItemRequest.Quantity
	err = db.Model(&model.CartItem{}).Create(&cartItem).Error
	if err != nil {
		return model.CartItem{}, err
	}
	return cartItem, nil
}

func ValidateProductById(productId string) (bool, error) {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Unable to connect User Service due to: " + err.Error())
	}
	defer conn.Close()
	client := productpb.NewProductServiceClient(conn)
	isProductValidRes, err := client.ValidateProduct(context.TODO(), &productpb.ValidateProductRequest{ProductId: productId})
	if err != nil && isProductValidRes == nil {
		return false, err
	}
	return true, nil
}
