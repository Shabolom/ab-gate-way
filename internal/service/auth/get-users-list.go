package auth

import (
	"context"
	authMicroserviceDto "gate-way/internal/dto/auth-microservice-dto"

	"go.uber.org/zap"
)

func (s *Service) GetUsersList(ctx context.Context) ([]*authMicroserviceDto.User, error) {
	s.logger.Info("get users list started")

	users, err := s.authAdapter.GetUsersList(ctx)
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
