package authService

import (
	"context"
	"gate-way/internal/dto"
	"gate-way/pkg/utils"

	"go.uber.org/zap"
)

func (s *Service) DeleteUser(ctx context.Context, tokens *dto.Tokens) (*dto.CommonResponse, error) {
	s.logger.Info("delete users started")

	if err := utils.ValidateTokens(tokens); err != nil {
		s.logger.Warn(
			"delete users validation failed",
			zap.Error(err),
		)
		return nil, err
	}

	response, err := s.authAdapter.DeleteUsers(ctx, tokens)
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
