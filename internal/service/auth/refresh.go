package auth

import (
	"context"
	"gate-way/internal/dto"

	"go.uber.org/zap"
)

func (s *Service) Refresh(ctx context.Context) (*dto.CommonResponse, error) {
	s.logger.Info("refresh started")

	response, err := s.authAdapter.Refresh(ctx)
	if err != nil {
		s.logger.Warn(
			"refresh failed",
			zap.Error(err),
		)
		return nil, err
	}

	s.logger.Info("refresh successful")

	return response, nil
}
