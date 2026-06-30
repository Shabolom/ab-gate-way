package ab

import (
	"context"
	abDto "gate-way/internal/dto/ab-dto"

	"go.uber.org/zap"
)

func (s *Service) GetCustomParams(ctx context.Context) (*abDto.GetCustomParamsResponse, error) {
	s.logger.Info("GetCustomParams started")

	response, err := s.abAdapter.GetCustomParams(ctx)
	if err != nil {
		s.logger.Warn("GetCustomParams failed", zap.Error(err))
		return nil, err
	}

	s.logger.Info("GetCustomParams finished")

	return response, nil
}
