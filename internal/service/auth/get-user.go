package auth

import (
	"context"
	authMicroserviceDto "gate-way/internal/dto/auth-microservice-dto"

	"go.uber.org/zap"
)

func (s *Service) GetUser(ctx context.Context) (*authMicroserviceDto.User, error) {
	s.logger.Info("get user started")

	user, err := s.authAdapter.GetUser(ctx)
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
