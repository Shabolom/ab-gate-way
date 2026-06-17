package ab

import (
	"context"
	"gate-way/internal/dto"
	"gate-way/pkg/shortcut"

	"go.uber.org/zap"
)

func (s *Service) SetReadyExperiment(ctx context.Context, experimentID int64) (*dto.CommonResponse, error) {
	s.logger.Info(
		"ab set ready experiment started",
		zap.Int64("experiment_id", experimentID),
	)

	if experimentID == 0 {
		s.logger.Warn("ab set ready experiment validation failed: experiment_id is empty")
		return nil, shortcut.ErrABExperimentIDRequired
	}

	response, err := s.abAdapter.SetReadyExperiment(ctx, experimentID)
	if err != nil {
		s.logger.Error(
			"ab set ready experiment failed",
			zap.Error(err),
		)
		return nil, err
	}

	s.logger.Info(
		"ab set ready experiment finished",
		zap.Int64("experiment_id", experimentID),
	)

	return response, nil
}
