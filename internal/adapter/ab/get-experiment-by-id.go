package abAdapter

import (
	"context"
	authv1 "gate-way/gen/proto"
	abDto "gate-way/internal/dto/ab-dto"
)

func (a *Adapter) GetExperimentByID(ctx context.Context, id int64) (*abDto.GetExperimentByIDResponse, error) {
	response, err := a.client.GetExperimentByID(ctx, &authv1.GetExperimentByIDRequest{Id: id})
	if err != nil {
		return nil, err
	}

	err = getExperimentErr(response.ErrInfoReason)
	if err != nil {
		return nil, err
	}

	responseExperiment := response.GetExperiment()

	var customGroups []abDto.GetParamGroup
	for _, customGroup := range responseExperiment.ParamsGroups {

		var params []abDto.GetCustomParamWithCondition
		for _, param := range customGroup.ParamsWithConditions {
			params = append(params, abDto.GetCustomParamWithCondition{
				ID:               param.GetId(),
				ParameterID:      param.GetParameterId(),
				ParameterGroupID: param.GetParameterGroupId(),
				Value:            param.GetValue(),
				Condition:        param.GetCondition(),
			})
		}

		customGroups = append(customGroups, abDto.GetParamGroup{
			ID:                   customGroup.Id,
			Percent:              customGroup.Percent,
			ParamsWithConditions: params,
		})
	}

	var groups []abDto.GetGroup
	for _, group := range responseExperiment.Groups {
		groups = append(groups, abDto.GetGroup{
			ID:                group.GetId(),
			Name:              group.GetName(),
			RollingPercentage: group.GetRollingPercentage(),
			DeviceID:          group.GetDeviceId(),
		})
	}

	return &abDto.GetExperimentByIDResponse{
		Message: response.Message,
		Experiment: &abDto.GetExperiment{
			ID:                responseExperiment.GetId(),
			Name:              responseExperiment.GetName(),
			Namespace:         responseExperiment.GetNamespace(),
			Status:            responseExperiment.GetStatus(),
			RolloutPercentage: responseExperiment.GetRolloutPercentage(),
			StartDate:         responseExperiment.GetStartDate().AsTime(),
			EndDate:           responseExperiment.GetEndDate().AsTime(),
			PassingCities:     responseExperiment.GetPassingCities(),
			ExcludedCities:    responseExperiment.GetExcludedCities(),
			PassingStores:     responseExperiment.GetPassingStores(),
			ExcludedStores:    responseExperiment.GetExcludedStores(),
			LayersID:          responseExperiment.GetLayersId(),
			ParamsGroups:      customGroups,
			Groups:            groups,
		},
	}, nil
}
