package abDto

type Layer struct {
	NamespaceID int64  `json:"namespace_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
