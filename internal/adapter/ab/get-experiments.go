package abAdapter

import (
	"context"
	abDto "gate-way/internal/dto/ab-dto"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *Adapter) GetExperiments(ctx context.Context) (*abDto.GetExperimentsResponse, error) {
	response, err := a.client.GetExperiments(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	err = getExperimentsErr(response.ErrInfoReason)
	if err != nil {
		return nil, err
	}

	responseExperiments := response.GetExperiments()

	var responseBodyExperiments []*abDto.GetExperiment
	for _, experiment := range responseExperiments {
		var experimentBody abDto.GetExperiment

		var groups []abDto.GetGroup
		for _, group := range experiment.GetGroups() {
			experimentBody.Groups = append(groups, abDto.GetGroup{
				ID:                group.GetId(),
				Name:              group.GetName(),
				RollingPercentage: group.GetRollingPercentage(),
				DeviceID:          group.GetDeviceId(),
			})
		}

		var customGroups []abDto.GetParamGroup
		for _, customGroup := range experiment.GetParamsGroups() {
			var params abDto.GetParamGroup

			params.ID = customGroup.GetId()
			params.Percent = customGroup.GetPercent()

			for _, param := range customGroup.GetParamsWithConditions() {
				params.ParamsWithConditions = append(params.ParamsWithConditions, abDto.GetCustomParamWithCondition{
					ID:               param.GetId(),
					ParameterID:      param.GetParameterId(),
					ParameterGroupID: param.GetParameterGroupId(),
					Value:            param.GetValue(),
					Condition:        param.GetCondition(),
				})
			}

			customGroups = append(customGroups, params)
		}

		responseBodyExperiments = append(responseBodyExperiments, &abDto.GetExperiment{
			ID:                experiment.GetId(),
			Name:              experiment.GetName(),
			Namespace:         experiment.GetNamespace(),
			Status:            experiment.GetStatus(),
			RolloutPercentage: experiment.GetRolloutPercentage(),
			StartDate:         experiment.GetStartDate().AsTime(),
			EndDate:           experiment.GetEndDate().AsTime(),
			PassingCities:     experiment.GetPassingCities(),
			ExcludedCities:    experiment.GetExcludedCities(),
			PassingStores:     experiment.GetPassingStores(),
			ExcludedStores:    experiment.GetExcludedStores(),
			LayersID:          experiment.GetLayersId(),
			ParamsGroups:      customGroups,
			Groups:            groups,
		})
	}

	return &abDto.GetExperimentsResponse{
		Message:     response.Message,
		Experiments: responseBodyExperiments,
	}, nil
}
