package abAdapter

import (
	"context"
	abDto "gate-way/internal/dto/ab-dto"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *Adapter) GetFeatureToggles(ctx context.Context) (*abDto.GetFeatureTogglesResponse, error) {
	response, err := a.client.GetFeatureToggles(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	err = getFeatureTogglesErr(response.ErrInfoReason)
	if err != nil {
		return nil, err
	}

	responseFeatureToggles := response.GetFeatureToggles()

	var featureTogglesBody []*abDto.GetFeatureToggle
	for _, feature := range responseFeatureToggles {
		featureToggle := &abDto.GetFeatureToggle{
			ID:                feature.GetId(),
			NamespaceID:       feature.GetNamespaceId(),
			Name:              feature.GetName(),
			Status:            feature.GetStatus(),
			RolloutPercentage: feature.RolloutPercentage,
			IOS:               feature.Ios,
			Android:           feature.Android,
			Web:               feature.Web,
			CreatedAt:         feature.CreatedAt.AsTime(),
		}

		if feature.UpdatedAt != nil {
			updatedAt := feature.UpdatedAt.AsTime()
			featureToggle.UpdatedAt = &updatedAt
		}

		if feature.DeletedAt != nil {
			deletedAt := feature.GetDeletedAt().AsTime()
			featureToggle.DeletedAt = &deletedAt
		}

		featureTogglesBody = append(featureTogglesBody, featureToggle)
	}

	return &abDto.GetFeatureTogglesResponse{
		Message:        response.Message,
		FeatureToggles: featureTogglesBody,
	}, nil
}
