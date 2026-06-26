package abDto

type UserFeatureReq struct {
	UserID    int64  `json:"user_id"`
	Namespace string `json:"namespace"`
	Platform  string `json:"platform"`
}
