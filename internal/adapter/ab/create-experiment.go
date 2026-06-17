package abAdapter

import (
	"context"
	authv1 "gate-way/gen/proto"
	"gate-way/internal/dto"
	abDto "gate-way/internal/dto/ab-dto"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (a *Adapter) CreateExperiment(ctx context.Context, request *abDto.CreateExperimentRequest) (*dto.CommonResponse, error) {
	requestBody := &authv1.CreateExperimentRequest{
		Name:              request.Name,
		RolloutPercentage: request.RolloutPercentage,
		StartDate:         timestamppb.New(request.StartDate),
		EndDate:           timestamppb.New(request.EndDate),
		LayersId:          request.LayersID,
		PassingCities:     request.PassingCities,
		ExcludedCities:    request.ExcludedCities,
		PassingStores:     request.PassingStores,
		ExcludedStores:    request.ExcludedStores,
	}

	for _, group := range request.Groups {
		requestBody.Groups = append(requestBody.Groups, &authv1.Group{
			Name:              group.Name,
			RollingPercentage: group.RollingPercentage,
			DeviceId:          group.DeviceID,
		})
	}

	for _, customParamGroup := range request.CustomParamGroups {
		var customParams []*authv1.CustomParam

		for _, customParam := range customParamGroup.CustomParam {
			customParamData := &authv1.CustomParam{
				ParametrId: customParam.ParamID,
				Condition:  customParam.Condition,
				Value:      customParam.Value,
			}
			customParams = append(customParams, customParamData)
		}

		requestBody.CustomParamGroups = append(requestBody.CustomParamGroups, &authv1.CustomParamGroup{
			Percentage:  customParamGroup.Percentage,
			CustomParam: customParams,
		})
	}

	response, err := a.client.CreateExperiment(ctx, requestBody)
	if err != nil {
		return nil, err
	}

	err = stockReplyErr(response.ErrInfoReason)
	if err != nil {
		return nil, err
	}

	return &dto.CommonResponse{Message: response.Message}, nil
}
