package ab

import (
	"context"
	abDto "gate-way/internal/dto/ab-dto"
	"gate-way/pkg/shortcut"

	"go.uber.org/zap"
)

func (s *Service) GetLayerByID(ctx context.Context, id int64) (*abDto.GetLayerByIDResponse, error) {
	if id == 0 {
		s.logger.Warn("GetLayerByID validate err, id is 0")
		return nil, shortcut.ErrValidation
	}

	s.logger.Info("GetLayerByID started",
		zap.Int64("id", id))

	response, err := s.abAdapter.GetLayerByID(ctx, id)
	if err != nil {
		s.logger.Warn("GetLayerByID failed", zap.Error(err))
		return nil, err
	}

	s.logger.Info("GetLayerByID finished",
		zap.Int64("id", id))

	return response, nil
}
