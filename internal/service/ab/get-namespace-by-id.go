package ab

import (
	"context"
	abDto "gate-way/internal/dto/ab-dto"
	"gate-way/pkg/shortcut"

	"go.uber.org/zap"
)

func (s *Service) GetNamespaceByID(ctx context.Context, id int64) (*abDto.GetNamespaceByIDResponse, error) {
	if id == 0 {
		s.logger.Warn("GetNamespaceByID validate err, id is 0")
		return nil, shortcut.ErrValidation
	}

	s.logger.Info("GetNamespaceByID started",
		zap.Int64("id", id))

	response, err := s.abAdapter.GetNamespaceByID(ctx, id)
	if err != nil {
		s.logger.Warn("GetNamespaceByID failed", zap.Error(err))
		return nil, err
	}

	s.logger.Info("GetNamespaceByID finished",
		zap.Int64("id", id))

	return response, nil
}
