package structs

type SuccessResponse struct {
	Status  string `json:"status"`
	Message any    `json:"message"`
}
