package auth

import (
	"context"
	"gate-way/internal/dto"
	authMicroserviceDto "gate-way/internal/dto/auth-microservice-dto"

	"go.uber.org/zap"
)

type authAdapter interface {
	Refresh(ctx context.Context) (*dto.CommonResponse, error)
	Register(ctx context.Context, register *authMicroserviceDto.Register) (*dto.CommonResponse, *dto.Tokens, error)
	Logout(ctx context.Context) (*dto.CommonResponse, error)
	Login(ctx context.Context, login *authMicroserviceDto.Login) (*dto.CommonResponse, *dto.Tokens, error)
	GetUsersList(ctx context.Context) ([]*authMicroserviceDto.User, error)
	GetUser(ctx context.Context) (*authMicroserviceDto.User, error)
	DeleteUsers(ctx context.Context) (*dto.CommonResponse, error)
	UpdateUsers(ctx context.Context, updateUser *authMicroserviceDto.UpdateUser) (*dto.CommonResponse, error)
	Checker(ctx context.Context) (context.Context, error)
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
