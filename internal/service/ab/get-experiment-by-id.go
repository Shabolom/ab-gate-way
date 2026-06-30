package ab

import (
	"context"
	abDto "gate-way/internal/dto/ab-dto"
	"gate-way/pkg/shortcut"

	"go.uber.org/zap"
)

func (s *Service) GetExperimentByID(ctx context.Context, id int64) (*abDto.GetExperimentByIDResponse, error) {
	if id == 0 {
		s.logger.Warn("GetExperimentByID validate err, id is 0")
		return nil, shortcut.ErrValidation
	}

	s.logger.Info("GetExperimentByID started",
		zap.Int64("id", id))

	response, err := s.abAdapter.GetExperimentByID(ctx, id)
	if err != nil {
		s.logger.Warn("GetExperimentByID failed", zap.Error(err))
		return nil, err
	}

	s.logger.Info("GetExperimentByID finished",
		zap.Int64("id", id))

	return response, nil
}
