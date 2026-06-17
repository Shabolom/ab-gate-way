package abAdapter

import (
	"context"
	authv1 "gate-way/gen/proto"
	"gate-way/internal/dto"
	abDto "gate-way/internal/dto/ab-dto"
)

func (a *Adapter) CreateNamespace(ctx context.Context, request *abDto.Namespace) (*dto.CommonResponse, error) {
	requestBody := &authv1.CreateNamespaceRequest{
		Name:        request.Name,
		Description: request.Description,
	}

	response, err := a.client.CreateNamespace(ctx, requestBody)
	if err != nil {
		return nil, err
	}

	err = stockReplyErr(response.ErrInfoReason)
	if err != nil {
		return nil, err
	}

	return &dto.CommonResponse{Message: response.Message}, nil
}
