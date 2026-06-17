package ab

import (
	"context"
	"gate-way/internal/dto"
	abDto "gate-way/internal/dto/ab-dto"
	"gate-way/pkg/shortcut"

	"go.uber.org/zap"
)

func (s *Service) CreateCustomParam(ctx context.Context, request *abDto.CreateCustomParamRequest) (*dto.CommonResponse, error) {
	s.logger.Info(
		"ab create custom param started",
		zap.String("name", request.Name),
		zap.Int64("namespace_id", request.NamespaceID),
		zap.String("type", request.Type),
	)

	switch {
	case request.Name == "":
		s.logger.Warn("ab create custom param validation failed: name is empty")
		return nil, shortcut.ErrABCustomParamNameRequired

	case request.NamespaceID == 0:
		s.logger.Warn("ab create custom param validation failed: namespace_id is empty")
		return nil, shortcut.ErrABCustomParamNamespaceRequired

	case request.Type == "":
		s.logger.Warn("ab create custom param validation failed: type is empty")
		return nil, shortcut.ErrABCustomParamTypeRequired
	}

	response, err := s.abAdapter.CreateCustomParam(ctx, request)
	if err != nil {
		s.logger.Error(
			"ab create custom param failed",
			zap.Error(err),
		)
		return nil, err
	}

	s.logger.Info(
		"ab create custom param finished",
		zap.String("name", request.Name),
	)

	return response, nil
}
