package ab

import (
	"context"
	abDto "gate-way/internal/dto/ab-dto"

	"go.uber.org/zap"
)

func (s *Service) GetLayers(ctx context.Context) (*abDto.GetLayersResponse, error) {
	s.logger.Info("GetLayers started")

	response, err := s.abAdapter.GetLayers(ctx)
	if err != nil {
		s.logger.Warn("GetLayers failed", zap.Error(err))
		return nil, err
	}

	s.logger.Info("GetLayers finished")

	return response, nil
}
