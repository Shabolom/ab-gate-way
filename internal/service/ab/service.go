package ab

import (
	"context"
	"gate-way/internal/dto"
	abDto "gate-way/internal/dto/ab-dto"

	"go.uber.org/zap"
)

type abAdapter interface {
	UserExperiment(ctx context.Context, request *abDto.UserExperimentRequest) (*abDto.ExperimentsReply, error)
	CreateExperiment(ctx context.Context, request *abDto.CreateExperimentRequest) (*dto.CommonResponse, error)
	CreateNamespace(ctx context.Context, request *abDto.Namespace) (*dto.CommonResponse, error)
	CreateLayer(ctx context.Context, layer *abDto.Layer) (*dto.CommonResponse, error)
	SetReadyExperiment(ctx context.Context, request int64) (*dto.CommonResponse, error)
	SetStopedExperiment(ctx context.Context, request int64) (*dto.CommonResponse, error)
	CreateCustomParam(ctx context.Context, request *abDto.CreateCustomParamRequest) (*dto.CommonResponse, error)
}

type Service struct {
	abAdapter abAdapter
	logger    *zap.Logger
}

func New(
	abAdapter abAdapter,
	logger *zap.Logger,
) *Service {
	return &Service{
		abAdapter: abAdapter,
		logger:    logger,
	}
}
