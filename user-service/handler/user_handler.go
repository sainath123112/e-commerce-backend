package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sainath123112/e-commerce-backend/user-service/model"
	"github.com/sainath123112/e-commerce-backend/user-service/pkg/jwt"
	"github.com/sainath123112/e-commerce-backend/user-service/service"
	"gorm.io/gorm"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var userRegisterRequestDto model.UserRegisterRequestDto
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&userRegisterRequestDto)
	isExist, _ := service.IsUserExist(userRegisterRequestDto.Email)

	if isExist {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&model.UserRegisterResponseDto{Message: "User Already Exists..! Try different email..!", Status: false})
		return
	}

	isRegistered, err := service.RegisterUserService(userRegisterRequestDto)
	if !isRegistered {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&model.UserRegisterResponseDto{Message: err.Error(), Status: false})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&model.UserRegisterResponseDto{Message: "User Registered Successfully..!", Status: true})

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var userLoginRequestDto model.UserLoginRequestDto
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&userLoginRequestDto)
	isExist, err := service.IsUserExist(userLoginRequestDto.Username)
	if !isExist {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&model.UserLoginResponseDto{Message: "No user found with username " + userLoginRequestDto.Username, Status: false, Token: err.Error()})
		return
	}
	userId, isAuthenticated, err := service.IsAuthenticated(userLoginRequestDto.Username, userLoginRequestDto.Password)
	if !isAuthenticated {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&model.UserLoginResponseDto{Message: "Invalid Password, Try again! ", Status: false, Token: err.Error()})
		return
	}
	token, err := jwt.GenerateToken(userLoginRequestDto.Username)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&model.UserLoginResponseDto{Message: "Unable to generate token", Status: false, Token: err.Error()})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&model.UserLoginResponseDto{UserId: userId, Message: "User successfully loggedin..!", Status: true, Token: token})
}

func GetUserDetails(w http.ResponseWriter, r *http.Request) {
	var userDetails model.UserDetails
	vars := mux.Vars(r)
	userid, err := uuid.Parse(vars["userid"])
	if err != nil {
		log.Fatalln("Error converting userid to uuid")
	}
	err = service.GetDetails(userid, &userDetails)
	if err == gorm.ErrRecordNotFound {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&model.ErrorResponse{Message: "token is not belongs to this user with user id: " + vars["userid"], ErrorString: "Token mismatch"})
		return
	}
	json.NewEncoder(w).Encode(&userDetails)

}
