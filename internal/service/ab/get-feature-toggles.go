package ab

import (
	"context"
	abDto "gate-way/internal/dto/ab-dto"

	"go.uber.org/zap"
)

func (s *Service) GetFeatureToggles(ctx context.Context) (*abDto.GetFeatureTogglesResponse, error) {
	s.logger.Info("GetFeatureToggles started")

	response, err := s.abAdapter.GetFeatureToggles(ctx)
	if err != nil {
		s.logger.Warn("GetFeatureToggles failed", zap.Error(err))
		return nil, err
	}

	s.logger.Info("GetFeatureToggles finished")

	return response, nil
}
