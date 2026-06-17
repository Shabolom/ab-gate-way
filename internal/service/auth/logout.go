package auth

import (
	"context"
	"gate-way/internal/dto"

	"go.uber.org/zap"
)

func (s *Service) Logout(ctx context.Context) (*dto.CommonResponse, error) {
	s.logger.Info("logout started")

	response, err := s.authAdapter.Logout(ctx)
	if err != nil {
		s.logger.Warn(
			"logout failed",
			zap.Error(err),
		)
		return nil, err
	}

	s.logger.Info("logout successful")

	return response, nil
}
