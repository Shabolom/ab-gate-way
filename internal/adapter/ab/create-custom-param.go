package abAdapter

import (
	"context"
	authv1 "gate-way/gen/proto"
	"gate-way/internal/dto"
	abDto "gate-way/internal/dto/ab-dto"
)

func (a *Adapter) CreateCustomParam(ctx context.Context, request *abDto.CreateCustomParamRequest) (*dto.CommonResponse, error) {
	requestBody := &authv1.CreateCustomParamRequest{
		Name:        request.Name,
		NamespaceId: request.NamespaceID,
		Type:        request.Type,
	}

	response, err := a.client.CreateCustomParam(ctx, requestBody)
	if err != nil {
		return nil, err
	}

	err = stockReplyErr(response.ErrInfoReason)
	if err != nil {
		return nil, err
	}

	return &dto.CommonResponse{Message: response.Message}, nil
}
