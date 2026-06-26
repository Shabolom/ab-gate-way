package ab

import (
	"context"
	"gate-way/internal/dto"
	abDto "gate-way/internal/dto/ab-dto"
	"gate-way/pkg/shortcut"

	"go.uber.org/zap"
)

func (s *Service) UpdateFeatureToggleRollout(ctx context.Context, request *abDto.UpdateFeature) (*dto.CommonResponse, error) {
	s.logger.Info(
		"update feature toggle rollout started",
		zap.Int64("feature_toggle_id", request.FeatureID),
	)

	err := updateFeatureValidation(request)
	if err != nil {
		s.logger.Warn(
			"update feature toggle rollout validation failed",
			zap.Int64("feature_toggle_id", request.FeatureID),
			zap.Error(err),
		)

		return nil, err
	}

	response, err := s.abAdapter.UpdateFeatureToggleRollout(ctx, request)
	if err != nil {
		s.logger.Error(
			"update feature toggle rollout failed",
			zap.Int64("feature_toggle_id", request.FeatureID),
			zap.Error(err),
		)

		return nil, err
	}

	s.logger.Info(
		"update feature toggle rollout finished",
		zap.Int64("feature_toggle_id", request.FeatureID),
	)

	return response, nil
}

func updateFeatureValidation(request *abDto.UpdateFeature) error {
	switch {
	case request.FeatureID == 0:
		return shortcut.ErrFeatureToggleID

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
