package authService

import (
	"context"
	"gate-way/internal/dto"
	authMicroserviceDto "gate-way/internal/dto/auth-microservice-dto"
	"gate-way/pkg/utils"

	"go.uber.org/zap"
)

func (s *Service) GetUser(ctx context.Context, tokens *dto.Tokens) (*authMicroserviceDto.User, error) {
	s.logger.Info("get user started")

	if err := utils.ValidateTokens(tokens); err != nil {
		s.logger.Warn(
			"get user validation failed",
			zap.Error(err),
		)
		return nil, err
	}

	user, err := s.authAdapter.GetUser(ctx, tokens)
	if err != nil {
		s.logger.Warn(
			"get user failed",
			zap.Error(err),
		)
		return nil, err
	}

	s.logger.Info(
		"get user successful",
		zap.String("user_id", user.ID),
		zap.String("mail", user.Mail),
	)

	return user, nil
}
