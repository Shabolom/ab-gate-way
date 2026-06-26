package abAdapter

import (
	"context"
	authv1 "gate-way/gen/proto"
	"gate-way/internal/dto"
	abDto "gate-way/internal/dto/ab-dto"
)

func (a *Adapter) CreateFeatureToggle(ctx context.Context, request *abDto.Feature) (*dto.CommonResponse, error) {
	req := &authv1.CreateFeatureToggleRequest{
		Name:                     request.Name,
		NamespaceId:              request.NamespaceID,
		IosRolloutPercentage:     request.IosRolloutPercentage,
		AndroidRolloutPercentage: request.AndroidRolloutPercentage,
		WebRolloutPercentage:     request.WebRolloutPercentage,
		RolloutPercentage:        request.RolloutPercentage,
	}

	response, err := a.client.CreateFeatureToggle(ctx, req)
	if err != nil {
		return nil, err
	}

	err = stockReplyErr(response.ErrInfoReason)
	if err != nil {
		return nil, err
	}

	return &dto.CommonResponse{Message: response.Message}, nil
}
