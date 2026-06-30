package ab

import (
	"context"
	abDto "gate-way/internal/dto/ab-dto"
	"gate-way/pkg/shortcut"

	"go.uber.org/zap"
)

func (s *Service) GetCustomParamByID(ctx context.Context, id int64) (*abDto.GetCustomParamByIDResponse, error) {
	if id == 0 {
		s.logger.Warn("GetCustomParamByID validate err, id is 0")
		return nil, shortcut.ErrValidation
	}

	s.logger.Info("GetCustomParamByID started",
		zap.Int64("id", id))

	response, err := s.abAdapter.GetCustomParamByID(ctx, id)
	if err != nil {
		s.logger.Warn("GetCustomParamByID failed", zap.Error(err))
		return nil, err
	}

	s.logger.Info("GetCustomParamByID finished",
		zap.Int64("id", id))

	return response, nil
}
