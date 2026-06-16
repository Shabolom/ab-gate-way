package user

import (
	"context"
	"gate-way/internal/dto"
	authMicroserviceDto "gate-way/internal/dto/auth-microservice-dto"
)

type authService interface {
	GetUsersList(ctx context.Context, tokens *dto.Tokens) ([]*authMicroserviceDto.User, error)
	GetUser(ctx context.Context, tokens *dto.Tokens) (*authMicroserviceDto.User, error)
	DeleteUser(ctx context.Context, tokens *dto.Tokens) (*dto.CommonResponse, error)
	UpdateUser(ctx context.Context, updateUser *authMicroserviceDto.UpdateUser, tokens *dto.Tokens) (*dto.CommonResponse, error)
}

type Handler struct {
	authService authService
}

func New(authService authService) *Handler {
	return &Handler{
		authService: authService,
	}
}
