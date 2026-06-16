package authMicroserviceDto

import "time"

type User struct {
	ID        string    `json:"id"`
	Mail      string    `json:"mail"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	CreatedAT time.Time `json:"created_at"`
	UpdatedAT time.Time `json:"updated_at"`
}
