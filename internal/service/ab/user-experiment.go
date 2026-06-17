package ab

import (
	"context"
	abDto "gate-way/internal/dto/ab-dto"
	"gate-way/pkg/shortcut"

	"go.uber.org/zap"
)

func (s *Service) UserExperiment(ctx context.Context, request *abDto.UserExperimentRequest) (*abDto.ExperimentsReply, error) {
	s.logger.Info(
		"ab user experiment started",
		zap.Int64("split_id", request.SplitID),
		zap.String("namespace", request.Namespace),
		zap.Int64("device_id", request.DeviceID),
		zap.String("city", request.City),
		zap.String("store", request.Store),
		zap.Int("params_count", len(request.Params)),
	)

	switch {
	case request.Namespace == "":
		s.logger.Warn("ab user experiment validation failed: namespace is empty")
		return nil, shortcut.ErrABNamespaceRequired

	case request.SplitID == 0:
		s.logger.Warn("ab user experiment validation failed: split_id is empty")
		return nil, shortcut.ErrABSplitIDRequired
	}

	response, err := s.abAdapter.UserExperiment(ctx, request)
	if err != nil {
		s.logger.Error(
			"ab user experiment failed",
			zap.Int64("split_id", request.SplitID),
			zap.String("namespace", request.Namespace),
			zap.Error(err),
		)
		return nil, err
	}

	s.logger.Info(
		"ab user experiment finished",
		zap.Int64("split_id", request.SplitID),
		zap.String("namespace", request.Namespace),
		zap.Int("experiments_count", len(response.ExperimentsReply)),
	)

	return response, nil
}
