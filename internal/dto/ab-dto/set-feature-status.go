package abDto

type SetFeatureStatus struct {
	FeatureID int64  `json:"feature_id,omitempty"`
	Status    string `json:"status,omitempty"`
}
