package experiments

import (
	"context"
	"gate-way/internal/dto"
	abDto "gate-way/internal/dto/ab-dto"
)

type abService interface {
	CreateExperiment(ctx context.Context, request *abDto.CreateExperimentRequest) (*dto.CommonResponse, error)
	SetStopedExperiment(ctx context.Context, experimentID int64) (*dto.CommonResponse, error)
	SetReadyExperiment(ctx context.Context, experimentID int64) (*dto.CommonResponse, error)
	UserExperiment(ctx context.Context, request *abDto.UserExperimentRequest) (*abDto.ExperimentsReply, error)
}

type Handler struct {
	abService abService
}

func New(abService abService) *Handler {
	return &Handler{
		abService: abService,
	}
}
