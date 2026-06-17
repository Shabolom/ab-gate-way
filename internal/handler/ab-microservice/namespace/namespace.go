package namespace

import (
	"context"
	"gate-way/internal/dto"
	abDto "gate-way/internal/dto/ab-dto"
)

type abService interface {
	CreateNamespace(ctx context.Context, request *abDto.Namespace) (*dto.CommonResponse, error)
}

type Handler struct {
	abService abService
}

func New(abService abService) *Handler {
	return &Handler{
		abService: abService,
	}
}
