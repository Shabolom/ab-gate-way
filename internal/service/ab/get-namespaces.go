package ab

import (
	"context"
	abDto "gate-way/internal/dto/ab-dto"

	"go.uber.org/zap"
)

func (s *Service) GetNamespaces(ctx context.Context) (*abDto.GetNamespacesResponse, error) {
	s.logger.Info("GetNamespaces started")

	response, err := s.abAdapter.GetNamespaces(ctx)
	if err != nil {
		s.logger.Warn("GetNamespaces failed", zap.Error(err))
		return nil, err
	}

	s.logger.Info("GetNamespaces finished")

	return response, nil
}
