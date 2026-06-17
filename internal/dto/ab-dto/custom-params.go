package abDto

type CreateCustomParamRequest struct {
	Name        string `json:"name"`
	NamespaceID int64  `json:"namespace_id"`
	Type        string `json:"type"`
}
