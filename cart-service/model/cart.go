package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserId    uuid.UUID  `gorm:"user_id" json:"user_id"`
	CartItems []CartItem `gorm:"foreignKey:CartId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"cart_items"`
}

type CartItem struct {
	gorm.Model
	CartId    uint      `gorm:"cart_id" json:"cart_id"`
	ProductId uuid.UUID `gorm:"product_id" json:"product_id"`
	Quantity  int       `gorm:"quantity" json:"quantity"`
}

type CartItemRequest struct {
	ProductId uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
}

type ErrorResponse struct {
	Message     string `json:"message"`
	ErrorString string `json:"error"`
}
