package auth

import (
	"context"
	"gate-way/internal/dto"
	authMicroserviceDto "gate-way/internal/dto/auth-microservice-dto"
	"gate-way/internal/service/auth"
)

type authService interface {
	Refresh(ctx context.Context) (*dto.CommonResponse, error)
	Login(ctx context.Context, login *authMicroserviceDto.Login) (*dto.CommonResponse, *dto.Tokens, error)
	Logout(ctx context.Context) (*dto.CommonResponse, error)
	Register(ctx context.Context, register *authMicroserviceDto.Register) (*dto.CommonResponse, *dto.Tokens, error)
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
