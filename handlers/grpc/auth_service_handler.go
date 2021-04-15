package grpc

import (
	"context"
	"log"

	"github.com/tuancamtbtx/learning-go/genproto/go/common/entities"
	service "github.com/tuancamtbtx/learning-go/genproto/go/common/services"
)

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (a *AuthHandler) GetMe(ctx context.Context, req *service.AuthServiceGetMeRequest) (*service.AuthServiceGetMeResponse, error) {
	log.Printf("Receive message body from client: %s", req.Token)

	if req.Token == "" {
		return &service.AuthServiceGetMeResponse{}, nil
	}
	user := &entities.User{
		Name: "tuancam",
	}
	return &service.AuthServiceGetMeResponse{
		Status:  true,
		Message: "Success",
		Data:    user,
	}, nil
}
