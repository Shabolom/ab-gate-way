package auth

import (
	"context"
	authv1 "gate-way/gen/proto"
	"gate-way/internal/dto"
	authMicroserviceDto "gate-way/internal/dto/auth-microservice-dto"

	"google.golang.org/grpc"
)

type authService interface {
	Register(ctx context.Context, register *authMicroserviceDto.Register) (*dto.CommonResponse, error)
	Login(ctx context.Context, login *authMicroserviceDto.Login) (*dto.CommonResponse, error)
	Logout(ctx context.Context, tokens *dto.Tokens) (*dto.CommonResponse, error)
	Refresh(ctx context.Context, tokens *dto.Tokens) (*dto.CommonResponse, error)
}

type Handler struct {
	client      authv1.AccountServiceClient
	authService authService
}

func New(authConn *grpc.ClientConn, authService authService) *Handler {
	return &Handler{
		client:      authv1.NewAccountServiceClient(authConn),
		authService: authService,
	}
}
