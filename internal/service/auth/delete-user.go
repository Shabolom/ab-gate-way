package auth

import (
	"context"
	"gate-way/internal/dto"

	"go.uber.org/zap"
)

func (s *Service) DeleteUser(ctx context.Context) (*dto.CommonResponse, error) {
	s.logger.Info("delete users started")

	response, err := s.authAdapter.DeleteUsers(ctx)
	if err != nil {
		s.logger.Warn(
			"delete users failed",
			zap.Error(err),
		)
		return nil, err
	}

	s.logger.Info("delete user successful")

	return response, nil
}
