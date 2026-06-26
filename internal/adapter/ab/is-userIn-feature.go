package abAdapter

import (
	"context"
	authv1 "gate-way/gen/proto"
	abDto "gate-way/internal/dto/ab-dto"
)

func (a *Adapter) IsUserInFeature(ctx context.Context, request *abDto.UserFeatureReq) ([]*abDto.FeatureReply, error) {
	requestBody := &authv1.IsUserInFeatureRequest{
		UserId:    request.UserID,
		Namespace: request.Namespace,
		Platform:  request.Platform,
	}

	response, err := a.client.IsUserInFeature(ctx, requestBody)
	if err != nil {
		return nil, err
	}

	err = isUserInFeatureReplyErr(response.ErrInfoReason)
	if err != nil {
		return nil, err
	}

	features := make([]*abDto.FeatureReply, 0)

	for _, feature := range response.GetFeatures() {
		features = append(features, &abDto.FeatureReply{
			FeatureID:   feature.GetFeatureId(),
			FeatureName: feature.GetFeatureName(),
		})
	}

	return features, nil
}
