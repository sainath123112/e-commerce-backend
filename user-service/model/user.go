package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID      `gorm:"type:uuid;primarykey;" json:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	FirstName    string         `gorm:"not null" json:"first_name"`
	LastName     string         `gorm:"not null" json:"last_name"`
	Email        string         `gorm:"unique;not null;" json:"email"`
	PasswordHash string         `gorm:"not null" json:"password_hash"`
	Address      UserAddress    `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;" json:"address"`
}

type UserAddress struct {
	gorm.Model
	UserId     uuid.UUID `gorm:"type:uuid;not null;" json:"user_id"`
	Address1   string    `gorm:"not null" json:"address_1"`
	Address2   string    `json:"address_2"`
	City       string    `gorm:"not null" json:"city"`
	State      string    `gorm:"not null" json:"state"`
	Country    string    `gorm:"not null" json:"country"`
	PostalCode string    `gorm:"not null" json:"postal_code"`
}

type UserLoginRequestDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegisterRequestDto struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserLoginResponseDto struct {
	UserId  uuid.UUID `json:"user_id"`
	Message string    `json:"message"`
	Status  bool      `json:"status"`
	Token   string    `json:"token"`
}

type UserRegisterResponseDto struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}
type UserDetails struct {
	UserId    uuid.UUID `json:"user_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
}
