package ab

import (
	"context"
	"gate-way/internal/dto"
	abDto "gate-way/internal/dto/ab-dto"
	"gate-way/pkg/shortcut"

	"go.uber.org/zap"
)

func (s *Service) CreateLayer(ctx context.Context, request *abDto.Layer) (*dto.CommonResponse, error) {
	s.logger.Info(
		"ab create layer started",
		zap.String("name", request.Name),
		zap.Int64("namespace_id", request.NamespaceID),
	)

	switch {
	case request.NamespaceID == 0:
		s.logger.Warn("ab create layer validation failed: namespace_id is empty")
		return nil, shortcut.ErrABLayerNamespaceRequired

	case request.Name == "":
		s.logger.Warn("ab create layer validation failed: name is empty")
		return nil, shortcut.ErrABLayerNameRequired
	}

	response, err := s.abAdapter.CreateLayer(ctx, request)
	if err != nil {
		s.logger.Error(
			"ab create layer failed",
			zap.Error(err),
		)
		return nil, err
	}

	s.logger.Info(
		"ab create layer finished",
		zap.String("name", request.Name),
	)

	return response, nil
}
