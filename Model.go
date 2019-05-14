package crema

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type GenericResponse struct {
	Response
	Data interface{} `json:"data"`
}
