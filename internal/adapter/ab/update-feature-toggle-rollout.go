package abAdapter

import (
	"context"
	authv1 "gate-way/gen/proto"
	"gate-way/internal/dto"
	abDto "gate-way/internal/dto/ab-dto"
)

func (a *Adapter) UpdateFeatureToggleRollout(ctx context.Context, request *abDto.UpdateFeature) (*dto.CommonResponse, error) {
	requestBody := &authv1.UpdateFeatureToggleRolloutRequest{
		FeatureToggleId:          request.FeatureID,
		IosRolloutPercentage:     request.IosRolloutPercentage,
		AndroidRolloutPercentage: request.AndroidRolloutPercentage,
		WebRolloutPercentage:     request.WebRolloutPercentage,
		RolloutPercentage:        request.RolloutPercentage,
	}

	response, err := a.client.UpdateFeatureToggleRollout(ctx, requestBody)
	if err != nil {
		return nil, err
	}

	err = stockReplyErr(response.ErrInfoReason)
	if err != nil {
		return nil, err
	}

	return &dto.CommonResponse{Message: response.Message}, nil
}
