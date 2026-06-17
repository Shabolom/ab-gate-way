package abDto

type CustomParam struct {
	ParamID   int64  `json:"param_id"`
	Condition string `json:"condition"`
	Value     string `json:"value"`
}
