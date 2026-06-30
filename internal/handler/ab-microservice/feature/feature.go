package feature

import (
	"context"
	"gate-way/internal/dto"
	abDto "gate-way/internal/dto/ab-dto"
)

type abService interface {
	CreateFeature(ctx context.Context, request *abDto.Feature) (*dto.CommonResponse, error)
	IsFeatureEnabled(ctx context.Context, id int64) (*dto.CommonResponse, error)
	IsUserInFeature(ctx context.Context, feature *abDto.UserFeatureReq) ([]*abDto.FeatureReply, error)
	SetFeatureToggleStatus(ctx context.Context, request *abDto.SetFeatureStatus) (*dto.CommonResponse, error)
	UpdateFeatureToggleRollout(ctx context.Context, request *abDto.UpdateFeature) (*dto.CommonResponse, error)
	GetFeatureToggleByID(ctx context.Context, id int64) (*abDto.GetFeatureToggleByIDResponse, error)
	GetFeatureToggles(ctx context.Context) (*abDto.GetFeatureTogglesResponse, error)
}

type Handler struct {
	abService abService
}

func New(abService abService) *Handler {
	return &Handler{
		abService: abService,
	}
}
