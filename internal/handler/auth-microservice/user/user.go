package user

import (
	"context"
	"gate-way/internal/dto"
	authMicroserviceDto "gate-way/internal/dto/auth-microservice-dto"
	"gate-way/internal/service/auth"
)

type authService interface {
	GetUsersList(ctx context.Context) ([]*authMicroserviceDto.User, error)
	GetUser(ctx context.Context) (*authMicroserviceDto.User, error)
	DeleteUser(ctx context.Context) (*dto.CommonResponse, error)
	UpdateUser(ctx context.Context, updateUser *authMicroserviceDto.UpdateUser) (*dto.CommonResponse, error)
}

type Handler struct {
	authService authService
	aa          auth.Service
}

func New(authService authService) *Handler {
	return &Handler{
		authService: authService,
	}
}
