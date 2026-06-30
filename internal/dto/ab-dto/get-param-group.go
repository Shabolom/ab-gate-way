package abDto
type GetParamGroup struct {
	ID                   int64                         `json:"id"`
	Percent              int64                         `json:"percent"`
	ParamsWithConditions []GetCustomParamWithCondition `json:"params_with_conditions"`
}