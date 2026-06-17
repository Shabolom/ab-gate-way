package ab

import (
	"context"
	"gate-way/internal/dto"
	abDto "gate-way/internal/dto/ab-dto"
	"gate-way/pkg/shortcut"

	"go.uber.org/zap"
)

func (s *Service) CreateNamespace(ctx context.Context, request *abDto.Namespace) (*dto.CommonResponse, error) {
	s.logger.Info(
		"ab create namespace started",
		zap.String("name", request.Name),
	)

	switch {
	case request.Name == "":
		s.logger.Warn("ab create namespace validation failed: name is empty")
		return nil, shortcut.ErrABNamespaceNameRequired
	}

	response, err := s.abAdapter.CreateNamespace(ctx, request)
	if err != nil {
		s.logger.Error(
			"ab create namespace failed",
			zap.String("name", request.Name),
			zap.Error(err),
		)
		return nil, err
	}

	s.logger.Info(
		"ab create namespace finished",
		zap.String("name", request.Name),
	)

	return response, nil
}
