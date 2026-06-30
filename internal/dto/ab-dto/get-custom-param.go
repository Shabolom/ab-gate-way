package abDto
type GetCustomParam struct {
	ID          int64  `json:"id"`
	NamespaceID int64  `json:"namespace_id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
}