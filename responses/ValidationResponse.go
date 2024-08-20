package responses

type ValidationResponse struct {
	Message string `json:"message"`
	Errors  error  `json:"errors"`
}
