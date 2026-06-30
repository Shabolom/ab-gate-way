package ab

import (
	"context"
	abDto "gate-way/internal/dto/ab-dto"
	"gate-way/pkg/shortcut"

	"go.uber.org/zap"
)

func (s *Service) GetFeatureToggleByID(ctx context.Context, id int64) (*abDto.GetFeatureToggleByIDResponse, error) {
	if id == 0 {
		s.logger.Warn("GetFeatureToggleByID validate err, id is 0")
		return nil, shortcut.ErrValidation
	}

	s.logger.Info("GetFeatureToggleByID started",
		zap.Int64("id", id))

	response, err := s.abAdapter.GetFeatureToggleByID(ctx, id)
	if err != nil {
		s.logger.Warn("GetFeatureToggleByID failed", zap.Error(err))
		return nil, err
	}

	s.logger.Info("GetFeatureToggleByID finished",
		zap.Int64("id", id))

	return response, nil
}
