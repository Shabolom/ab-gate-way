package abDto

type GetExperimentByIDResponse struct {
	Message    string        `json:"message"`
	Experiment *GetExperiment `json:"experiment"`
}

type GetNamespaceByIDResponse struct {
	Message   string       `json:"message"`
	Namespace *GetNamespace `json:"namespace"`
}

type GetLayerByIDResponse struct {
	Message string   `json:"message"`
	Layer   *GetLayer `json:"layer"`
}

type GetCustomParamByIDResponse struct {
	Message     string         `json:"message"`
	CustomParam *GetCustomParam `json:"custom_param"`
}

type GetFeatureToggleByIDResponse struct {
	Message       string           `json:"message"`
	FeatureToggle *GetFeatureToggle `json:"feature_toggle"`
}

type GetExperimentsResponse struct {
	Message     string          `json:"message"`
	Experiments []*GetExperiment `json:"experiments"`
}

type GetNamespacesResponse struct {
	Message    string         `json:"message"`
	Namespaces []*GetNamespace `json:"namespaces"`
}

type GetLayersResponse struct {
	Message string     `json:"message"`
	Layers  []*GetLayer `json:"layers"`
}

type GetCustomParamsResponse struct {
	Message      string           `json:"message"`
	CustomParams []*GetCustomParam `json:"custom_params"`
}

type GetFeatureTogglesResponse struct {
	Message        string             `json:"message"`
	FeatureToggles []*GetFeatureToggle `json:"feature_toggles"`
}















