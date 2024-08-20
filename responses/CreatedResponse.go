package responses

type CreatedResponse struct {
	Data struct {
		ID int64 `json:"id"`
	} `json:"data"`
	Message string `json:"message"`
}
