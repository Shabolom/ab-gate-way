package ab

import (
	"context"
	abDto "gate-way/internal/dto/ab-dto"

	"go.uber.org/zap"
)

func (s *Service) GetExperiments(ctx context.Context) (*abDto.GetExperimentsResponse, error) {
	s.logger.Info("GetExperiments started")

	response, err := s.abAdapter.GetExperiments(ctx)
	if err != nil {
		s.logger.Warn("GetExperiments failed", zap.Error(err))
		return nil, err
	}

	s.logger.Info("GetExperiments finished")

	return response, nil
}
