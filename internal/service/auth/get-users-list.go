package authService

import (
	"context"
	"gate-way/internal/dto"
	authMicroserviceDto "gate-way/internal/dto/auth-microservice-dto"
	"gate-way/pkg/utils"

	"go.uber.org/zap"
)

func (s *Service) GetUsersList(ctx context.Context, tokens *dto.Tokens) ([]*authMicroserviceDto.User, error) {
	s.logger.Info("get users list started")

	if err := utils.ValidateTokens(tokens); err != nil {
		s.logger.Warn(
			"get users list validation failed",
			zap.Error(err),
		)
		return nil, err
	}

	users, err := s.authAdapter.GetUsersList(ctx, tokens)
	if err != nil {
		s.logger.Warn(
			"get users list failed",
			zap.Error(err),
		)
		return nil, err
	}

	s.logger.Info(
		"get users list successful",
		zap.Int("users_count", len(users)),
	)

	return users, nil
}
