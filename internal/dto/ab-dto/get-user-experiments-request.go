package abDto

type UserExperimentRequest struct {
	SplitID   int64  `json:"split_id"`
	Namespace string `json:"namespace"`

	City     string  `json:"city,omitempty"`
	Store    string  `json:"store,omitempty"`
	DeviceID int64   `json:"device_id,omitempty"`
	Params   []Param `json:"params,omitempty"`
}
