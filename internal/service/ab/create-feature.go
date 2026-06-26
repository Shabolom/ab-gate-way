package ab

import (
	"context"
	"gate-way/internal/dto"
	abDto "gate-way/internal/dto/ab-dto"
	"gate-way/pkg/shortcut"

	"go.uber.org/zap"
)

func (s *Service) CreateFeature(ctx context.Context, request *abDto.Feature) (*dto.CommonResponse, error) {
	s.logger.Info(
		"create feature toggle started",
		zap.String("name", request.Name),
		zap.Int64("namespace_id", request.NamespaceID),
	)

	err := createFeatureValidation(request)
	if err != nil {
		s.logger.Warn(
			"create feature toggle validation failed",
			zap.String("name", request.Name),
			zap.Int64("namespace_id", request.NamespaceID),
			zap.Error(err),
		)

		return nil, err
	}

	response, err := s.abAdapter.CreateFeatureToggle(ctx, request)
	if err != nil {
		s.logger.Error(
			"create feature toggle failed",
			zap.String("name", request.Name),
			zap.Int64("namespace_id", request.NamespaceID),
			zap.Error(err),
		)

		return nil, err
	}

	s.logger.Info(
		"create feature toggle finished",
		zap.String("name", request.Name),
		zap.Int64("namespace_id", request.NamespaceID),
	)

	return response, nil
}

func createFeatureValidation(request *abDto.Feature) error {
	switch {
	case request.Name == "":
		return shortcut.ErrFeatureToggleNameRequired

	case request.NamespaceID <= 0:
		return shortcut.ErrFeatureToggleNamespaceIDRequired

	case request.RolloutPercentage != nil &&
		(*request.RolloutPercentage < 0 || *request.RolloutPercentage > 100):
		return shortcut.ErrFeatureToggleRolloutOutOfRange

	case request.IosRolloutPercentage != nil &&
		(*request.IosRolloutPercentage < 0 || *request.IosRolloutPercentage > 100):
		return shortcut.ErrFeatureToggleIOSRolloutOutOfRange

	case request.AndroidRolloutPercentage != nil &&
		(*request.AndroidRolloutPercentage < 0 || *request.AndroidRolloutPercentage > 100):
		return shortcut.ErrFeatureToggleAndroidRolloutOutOfRange

	case request.WebRolloutPercentage != nil &&
		(*request.WebRolloutPercentage < 0 || *request.WebRolloutPercentage > 100):
		return shortcut.ErrFeatureToggleWebRolloutOutOfRange
	}

	return nil
}
