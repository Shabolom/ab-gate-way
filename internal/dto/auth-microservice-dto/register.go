package authMicroserviceDto

type Register struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
}
