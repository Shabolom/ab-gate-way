package ab

import (
	"context"
	"gate-way/internal/dto"
	abDto "gate-way/internal/dto/ab-dto"
	"gate-way/pkg/shortcut"

	"go.uber.org/zap"
)

func (s *Service) SetFeatureToggleStatus(ctx context.Context, request *abDto.SetFeatureStatus) (*dto.CommonResponse, error) {
	s.logger.Info(
		"set feature toggle status started",
		zap.Int64("feature_toggle_id", request.FeatureID),
		zap.String("status", request.Status),
	)

	switch {
	case request.FeatureID == 0:
		s.logger.Warn("set feature toggle status validation failed", zap.Error(shortcut.ErrFeatureToggleID))
		return nil, shortcut.ErrFeatureToggleID

	case request.Status == "":
		s.logger.Warn("set feature toggle status validation failed", zap.Error(shortcut.ErrFeatureStatusMissing))
		return nil, shortcut.ErrFeatureStatusMissing
	}

	response, err := s.abAdapter.SetFeatureToggleStatus(ctx, request)
	if err != nil {
		s.logger.Error(
			"set feature toggle status failed",
			zap.Int64("feature_toggle_id", request.FeatureID),
			zap.String("status", request.Status),
			zap.Error(err),
		)

		return nil, err
	}

	s.logger.Info(
		"set feature toggle status finished",
		zap.Int64("feature_toggle_id", request.FeatureID),
		zap.String("status", request.Status),
	)

	return response, nil
}
