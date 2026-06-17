package abDto

type Group struct {
	Name              string  `json:"name"`
	RollingPercentage int64   `json:"rolling_percentage"`
	DeviceID          []int64 `json:"device_id"`
}
