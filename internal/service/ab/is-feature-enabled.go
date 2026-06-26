package ab

import (
	"context"
	"gate-way/internal/dto"
	"gate-way/pkg/shortcut"

	"go.uber.org/zap"
)

func (s *Service) IsFeatureEnabled(ctx context.Context, id int64) (*dto.CommonResponse, error) {
	s.logger.Info("is feature enabled started", zap.Int64("feature_toggle_id", id))

	if id == 0 {
		s.logger.Warn("is feature enabled validation failed", zap.Error(shortcut.ErrFeatureToggleID))
		return nil, shortcut.ErrFeatureToggleID
	}

	response, err := s.abAdapter.IsFeatureEnabled(ctx, id)
	if err != nil {
		s.logger.Error(
			"is feature enabled failed",
			zap.Int64("feature_toggle_id", id),
			zap.Error(err),
		)

		return nil, err
	}

	s.logger.Info("is feature enabled finished", zap.Int64("feature_toggle_id", id))

	return response, nil
}
