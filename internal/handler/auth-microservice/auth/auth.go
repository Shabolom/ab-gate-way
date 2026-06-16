package auth

import (
	"context"
	"gate-way/internal/dto"
	authMicroserviceDto "gate-way/internal/dto/auth-microservice-dto"
)

type authService interface {
	Register(ctx context.Context, register *authMicroserviceDto.Register) (*dto.CommonResponse, error)
	Login(ctx context.Context, login *authMicroserviceDto.Login) (*dto.CommonResponse, error)
	Logout(ctx context.Context, tokens *dto.Tokens) (*dto.CommonResponse, error)
	Refresh(ctx context.Context, tokens *dto.Tokens) (*dto.CommonResponse, error)
}

type Handler struct {
	authService authService
}

func New(authService authService) *Handler {
	return &Handler{
		authService: authService,
	}
}
