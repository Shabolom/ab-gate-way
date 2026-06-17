package abDto

import "time"

type CreateExperimentRequest struct {
	Name              string    `json:"name"`
	RolloutPercentage int64     `json:"rollout_percentage"`
	StartDate         time.Time `json:"start_date"`
	EndDate           time.Time `json:"end_date"`
	LayersID          []int64   `json:"layers_id"`

	PassingCities     []string           `json:"passing_cities,omitempty"`
	ExcludedCities    []string           `json:"excluded_cities,omitempty"`
	PassingStores     []string           `json:"passing_stores,omitempty"`
	ExcludedStores    []string           `json:"excluded_stores,omitempty"`
	CustomParamGroups []CustomParamGroup `json:"custom_param_groups,omitempty"`

	Groups []Group `json:"groups"`
}
