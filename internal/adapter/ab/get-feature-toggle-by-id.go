package abAdapter

import (
	"context"
	authv1 "gate-way/gen/proto"
	abDto "gate-way/internal/dto/ab-dto"
)

func (a *Adapter) GetFeatureToggleByID(ctx context.Context, id int64) (*abDto.GetFeatureToggleByIDResponse, error) {
	response, err := a.client.GetFeatureToggleByID(ctx, &authv1.GetFeatureToggleByIDRequest{Id: id})
	if err != nil {
		return nil, err
	}

	err = getFeatureToggleErr(response.ErrInfoReason)
	if err != nil {
		return nil, err
	}

	responseFeatureToggle := response.GetFeatureToggle()

	responseFeatureToggleBody := &abDto.GetFeatureToggleByIDResponse{
		Message: response.Message,
		FeatureToggle: &abDto.GetFeatureToggle{
			ID:                responseFeatureToggle.GetId(),
			NamespaceID:       responseFeatureToggle.GetNamespaceId(),
			Name:              responseFeatureToggle.GetName(),
			Status:            responseFeatureToggle.GetStatus(),
			RolloutPercentage: responseFeatureToggle.RolloutPercentage,
			IOS:               responseFeatureToggle.Ios,
			Android:           responseFeatureToggle.Android,
			Web:               responseFeatureToggle.Web,
			CreatedAt:         responseFeatureToggle.CreatedAt.AsTime(),
		},
	}

	if responseFeatureToggle.CreatedAt != nil {
		createdAt := responseFeatureToggle.CreatedAt.AsTime()
		responseFeatureToggleBody.FeatureToggle.CreatedAt = createdAt
	}

	if responseFeatureToggle.UpdatedAt != nil {
		updatedAt := responseFeatureToggle.UpdatedAt.AsTime()
		responseFeatureToggleBody.FeatureToggle.UpdatedAt = &updatedAt
	}

	if responseFeatureToggle.DeletedAt != nil {
		deletedAt := responseFeatureToggle.DeletedAt.AsTime()
		responseFeatureToggleBody.FeatureToggle.DeletedAt = &deletedAt
	}

	return responseFeatureToggleBody, nil
}
