package authService

import (
	"context"
	"gate-way/internal/dto"
	"gate-way/pkg/utils"

	"go.uber.org/zap"
)

func (s *Service) Refresh(ctx context.Context, tokens *dto.Tokens) (*dto.CommonResponse, error) {
	s.logger.Info("refresh started")

	if err := utils.ValidateTokens(tokens); err != nil {
		s.logger.Warn(
			"refresh validation failed",
			zap.Error(err),
		)
		return nil, err
	}

	response, err := s.authAdapter.Refresh(ctx, tokens)
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
