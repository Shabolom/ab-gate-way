package ab

import (
	"context"
	abDto "gate-way/internal/dto/ab-dto"
	"gate-way/pkg/shortcut"

	"go.uber.org/zap"
)

func (s *Service) IsUserInFeature(ctx context.Context, feature *abDto.UserFeatureReq) ([]*abDto.FeatureReply, error) {
	s.logger.Info(
		"is user in feature started",
		zap.Int64("user_id", feature.UserID),
		zap.String("namespace", feature.Namespace),
		zap.String("platform", feature.Platform),
	)

	switch {
	case feature.UserID == 0:
		s.logger.Warn("is user in feature validation failed", zap.Error(shortcut.ErrFeatureToggleID))
		return nil, shortcut.ErrFeatureToggleID

	case feature.Namespace == "":
		s.logger.Warn("is user in feature validation failed", zap.Error(shortcut.ErrFeatureNamespaceMissing))
		return nil, shortcut.ErrFeatureNamespaceMissing
	}

	response, err := s.abAdapter.IsUserInFeature(ctx, feature)
	if err != nil {
		s.logger.Error(
			"is user in feature failed",
			zap.Int64("user_id", feature.UserID),
			zap.String("namespace", feature.Namespace),
			zap.String("platform", feature.Platform),
			zap.Error(err),
		)

		return nil, err
	}

	s.logger.Info(
		"is user in feature finished",
		zap.Int64("user_id", feature.UserID),
		zap.String("namespace", feature.Namespace),
		zap.String("platform", feature.Platform),
		zap.Int("features_count", len(response)),
	)

	return response, nil
}
