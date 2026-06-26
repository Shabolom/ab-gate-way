package abDto

type Feature struct {
	Name                     string `json:"name"`
	NamespaceID              int64  `json:"namespace_id"`
	IosRolloutPercentage     *int64 `json:"ios_rollout_percentage"`
	AndroidRolloutPercentage *int64 `json:"android_rollout_percentage"`
	WebRolloutPercentage     *int64 `json:"web_rollout_percentage"`
	RolloutPercentage        *int64 `json:"rollout_percentage"`
}
