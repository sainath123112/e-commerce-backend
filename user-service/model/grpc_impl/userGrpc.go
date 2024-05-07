package grpc_impl

import (
	"context"

	"github.com/google/uuid"
	"github.com/sainath123112/e-commerce-backend/user-service/pkg/jwt"
	pb "github.com/sainath123112/e-commerce-backend/user-service/proto"
	"github.com/sainath123112/e-commerce-backend/user-service/service"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserServiceServer) GetUserExist(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	isExist, err := service.IsUserExistById(in.Id)
	return &pb.GetUserResponse{IsExist: isExist}, err
}

func (s *UserServiceServer) ValidateToken(ctx context.Context, in *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	isTokenValid, err := jwt.ValidateToken(in.Token)
	if isTokenValid {
		return &pb.ValidateTokenResponse{IsTokenValid: true}, err
	}
	return &pb.ValidateTokenResponse{IsTokenValid: false}, err
}

func (s *UserServiceServer) GetUserNameFromToken(ctx context.Context, in *pb.UserNameFromTokenRequest) (*pb.UserNameFromTokenResponse, error) {
	emailFromToken, err := jwt.GetUsername(in.Token)
	return &pb.UserNameFromTokenResponse{EmailFromToken: emailFromToken}, err
}

func (s *UserServiceServer) GetUserNameFromUserId(ctx context.Context, in *pb.UserNameFromUserIdRequest) (*pb.UserNameFromUserIdResponse, error) {
	userId, _ := uuid.Parse(in.UserId)
	emailFromUserId, err := service.GetUserEmail(userId)
	return &pb.UserNameFromUserIdResponse{EmailFromUserId: emailFromUserId}, err
}
