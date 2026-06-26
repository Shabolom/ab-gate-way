package abDto

type UpdateFeature struct {
	FeatureID                int64  `json:"feature_id"`
	IosRolloutPercentage     *int64 `json:"ios_rollout_percentage"`
	AndroidRolloutPercentage *int64 `json:"android_rollout_percentage"`
	WebRolloutPercentage     *int64 `json:"web_rollout_percentage"`
	RolloutPercentage        *int64 `json:"rollout_percentage"`
}
