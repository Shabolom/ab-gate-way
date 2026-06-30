package abDto

import "time"

type GetFeatureToggle struct {
	ID                int64      `json:"id"`
	NamespaceID       int64      `json:"namespace_id"`
	Name              string     `json:"name"`
	Status            string     `json:"status"`
	RolloutPercentage *int64     `json:"rollout_percentage,omitempty"`
	IOS               *int64     `json:"ios,omitempty"`
	Android           *int64     `json:"android,omitempty"`
	Web               *int64     `json:"web,omitempty"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         *time.Time `json:"updated_at"`
	DeletedAt         *time.Time `json:"deleted_at,omitempty"`
}
