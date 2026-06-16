package authService

import (
	"context"
	"gate-way/internal/dto"
	authMicroserviceDto "gate-way/internal/dto/auth-microservice-dto"

	"go.uber.org/zap"
)

type authAdapter interface {
	Register(ctx context.Context, register *authMicroserviceDto.Register) (*dto.CommonResponse, error)
	Login(ctx context.Context, login *authMicroserviceDto.Login) (*dto.CommonResponse, error)
	Logout(ctx context.Context, tokens *dto.Tokens) (*dto.CommonResponse, error)
	Refresh(ctx context.Context, tokens *dto.Tokens) (*dto.CommonResponse, error)
	GetUsersList(ctx context.Context, tokens *dto.Tokens) ([]*authMicroserviceDto.User, error)
	GetUser(ctx context.Context, tokens *dto.Tokens) (*authMicroserviceDto.User, error)
	DeleteUsers(ctx context.Context, tokens *dto.Tokens) (*dto.CommonResponse, error)
	UpdateUsers(ctx context.Context, updateUser *authMicroserviceDto.UpdateUser, tokens *dto.Tokens) (*dto.CommonResponse, error)
}

type Service struct {
	authAdapter authAdapter
	logger      *zap.Logger
}

func New(authAdapter authAdapter, logger *zap.Logger) *Service {
	return &Service{
		authAdapter: authAdapter,
		logger:      logger,
	}
}
