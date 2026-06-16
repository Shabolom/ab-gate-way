package authService

import (
	"context"
	"gate-way/internal/dto"
	"gate-way/pkg/utils"

	"go.uber.org/zap"
)

func (s *Service) Logout(ctx context.Context, tokens *dto.Tokens) (*dto.CommonResponse, error) {
	s.logger.Info("logout started")

	if err := utils.ValidateTokens(tokens); err != nil {
		s.logger.Warn(
			"logout validation failed",
			zap.Error(err),
		)
		return nil, err
	}

	response, err := s.authAdapter.Logout(ctx, tokens)
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
