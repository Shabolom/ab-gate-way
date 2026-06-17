package abDto

type CustomParamGroup struct {
	Percentage  int64         `json:"percentage"`
	CustomParam []CustomParam `json:"custom_param"`
}
