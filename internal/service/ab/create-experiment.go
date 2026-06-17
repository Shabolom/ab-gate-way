package ab

import (
	"context"
	"gate-way/internal/dto"
	abDto "gate-way/internal/dto/ab-dto"
	"gate-way/pkg/shortcut"

	"go.uber.org/zap"
)

func (s *Service) CreateExperiment(ctx context.Context, request *abDto.CreateExperimentRequest) (*dto.CommonResponse, error) {
	s.logger.Info(
		"ab create experiment started",
		zap.String("name", request.Name),
		zap.Int64("rollout_percentage", request.RolloutPercentage),
		zap.Int("layers_count", len(request.LayersID)),
		zap.Int("groups_count", len(request.Groups)),
		zap.Int("custom_param_groups_count", len(request.CustomParamGroups)),
	)

	switch {
	case request.Name == "":
		s.logger.Warn("ab create experiment validation failed: name is empty")
		return nil, shortcut.ErrABExperimentNameRequired

	case len(request.LayersID) == 0:
		s.logger.Warn("ab create experiment validation failed: layers_id is empty")
		return nil, shortcut.ErrABExperimentLayersRequired

	case len(request.Groups) < 2:
		s.logger.Warn(
			"ab create experiment validation failed: groups count less than 2",
			zap.Int("groups_count", len(request.Groups)),
		)
		return nil, shortcut.ErrABExperimentGroupsMinCount

	case request.EndDate.Before(request.StartDate):
		s.logger.Warn(
			"ab create experiment validation failed: end_date before start_date",
			zap.Time("start_date", request.StartDate),
			zap.Time("end_date", request.EndDate),
		)
		return nil, shortcut.ErrABExperimentStartDateAfterEnd

	case request.RolloutPercentage < 0 || request.RolloutPercentage > 100:
		s.logger.Warn(
			"ab create experiment validation failed: rollout percentage out of range",
			zap.Int64("rollout_percentage", request.RolloutPercentage),
		)
		return nil, shortcut.ErrABExperimentRolloutOutOfRange
	}

	response, err := s.abAdapter.CreateExperiment(ctx, request)
	if err != nil {
		s.logger.Error(
			"ab create experiment failed",
			zap.String("name", request.Name),
			zap.Error(err),
		)
		return nil, err
	}

	s.logger.Info(
		"ab create experiment finished",
		zap.String("name", request.Name),
		zap.String("message", response.Message),
	)

	return response, nil
}
