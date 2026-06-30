package abDto
type GetLayer struct {
	ID          int64  `json:"id"`
	NamespaceID int64  `json:"namespace_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}