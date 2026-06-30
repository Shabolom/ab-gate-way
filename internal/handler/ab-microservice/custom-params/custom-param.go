package customParams

import (
	"context"
	"gate-way/internal/dto"
	abDto "gate-way/internal/dto/ab-dto"
)

type abService interface {
	CreateCustomParam(ctx context.Context, request *abDto.CreateCustomParamRequest) (*dto.CommonResponse, error)
	GetCustomParamByID(ctx context.Context, id int64) (*abDto.GetCustomParamByIDResponse, error)
	GetCustomParams(ctx context.Context) (*abDto.GetCustomParamsResponse, error)
}

type Handler struct {
	abService abService
}

func New(abService abService) *Handler {
	return &Handler{
		abService: abService,
	}
}
