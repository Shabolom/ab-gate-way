package abDto
type GetCustomParamWithCondition struct {
	ID               int64  `json:"id"`
	ParameterID      int64  `json:"parameter_id"`
	ParameterGroupID int64  `json:"parameter_group_id"`
	Value            string `json:"value"`
	Condition        string `json:"condition"`
}