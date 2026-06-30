package abDto
type GetGroup struct {
	ID                int64   `json:"id"`
	Name              string  `json:"name"`
	RollingPercentage int64   `json:"rolling_percentage"`
	DeviceID          []int64 `json:"device_id"`
}
