package abAdapter

import (
	"context"
	authv1 "gate-way/gen/proto"
	"gate-way/internal/dto"
	abDto "gate-way/internal/dto/ab-dto"
)

func (a *Adapter) SetFeatureToggleStatus(ctx context.Context, request *abDto.SetFeatureStatus) (*dto.CommonResponse, error) {
	requestBody := &authv1.SetFeatureToggleStatusRequest{
		FeatureToggleId: request.FeatureID,
		Status:          request.Status,
	}

	response, err := a.client.SetFeatureToggleStatus(ctx, requestBody)
	if err != nil {
		return nil, err
	}

	err = stockReplyErr(response.ErrInfoReason)
	if err != nil {
		return nil, err
	}

	return &dto.CommonResponse{Message: response.Message}, nil
}
