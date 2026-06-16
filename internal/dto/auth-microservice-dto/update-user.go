package authMicroserviceDto

type UpdateUser struct {
	Mail string `json:"mail"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
