package ab

import (
	"context"
	"gate-way/internal/dto"
	abDto "gate-way/internal/dto/ab-dto"

	"go.uber.org/zap"
)

type abAdapter interface {
	CreateNamespace(ctx context.Context, request *abDto.Namespace) (*dto.CommonResponse, error)

	CreateLayer(ctx context.Context, layer *abDto.Layer) (*dto.CommonResponse, error)

	CreateCustomParam(ctx context.Context, request *abDto.CreateCustomParamRequest) (*dto.CommonResponse, error)

	UserExperiment(ctx context.Context, request *abDto.UserExperimentRequest) (*abDto.ExperimentsReply, error)
	CreateExperiment(ctx context.Context, request *abDto.CreateExperimentRequest) (*dto.CommonResponse, error)
	SetReadyExperiment(ctx context.Context, request int64) (*dto.CommonResponse, error)
	SetStopedExperiment(ctx context.Context, request int64) (*dto.CommonResponse, error)

	CreateFeatureToggle(ctx context.Context, request *abDto.Feature) (*dto.CommonResponse, error)
	UpdateFeatureToggleRollout(ctx context.Context, request *abDto.UpdateFeature) (*dto.CommonResponse, error)
	SetFeatureToggleStatus(ctx context.Context, request *abDto.SetFeatureStatus) (*dto.CommonResponse, error)
	IsFeatureEnabled(ctx context.Context, id int64) (*dto.CommonResponse, error)
	IsUserInFeature(ctx context.Context, request *abDto.UserFeatureReq) ([]*abDto.FeatureReply, error)

	GetExperimentByID(ctx context.Context, id int64) (*abDto.GetExperimentByIDResponse, error)
	GetNamespaceByID(ctx context.Context, id int64) (*abDto.GetNamespaceByIDResponse, error)
	GetLayerByID(ctx context.Context, id int64) (*abDto.GetLayerByIDResponse, error)
	GetCustomParamByID(ctx context.Context, id int64) (*abDto.GetCustomParamByIDResponse, error)
	GetFeatureToggleByID(ctx context.Context, id int64) (*abDto.GetFeatureToggleByIDResponse, error)
	GetExperiments(ctx context.Context) (*abDto.GetExperimentsResponse, error)
	GetNamespaces(ctx context.Context) (*abDto.GetNamespacesResponse, error)
	GetLayers(ctx context.Context) (*abDto.GetLayersResponse, error)
	GetCustomParams(ctx context.Context) (*abDto.GetCustomParamsResponse, error)
	GetFeatureToggles(ctx context.Context) (*abDto.GetFeatureTogglesResponse, error)
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
