package abDto

import "time"

type GetExperiment struct {
	ID                int64           `json:"id"`
	Name              string          `json:"name"`
	Namespace         string          `json:"namespace"`
	Status            string          `json:"status"`
	RolloutPercentage int64           `json:"rollout_percentage"`
	StartDate         time.Time       `json:"start_date"`
	EndDate           time.Time       `json:"end_date"`
	PassingCities     []string        `json:"passing_cities"`
	ExcludedCities    []string        `json:"excluded_cities"`
	PassingStores     []string        `json:"passing_stores"`
	ExcludedStores    []string        `json:"excluded_stores"`
	LayersID          []int64         `json:"layers_id"`
	ParamsGroups      []GetParamGroup `json:"params_groups"`
	Groups            []GetGroup      `json:"groups"`
}
