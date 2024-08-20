package responses

import "ashleycoles/logbook-api/models"

type LogResponse struct {
	Data    []models.Log `json:"data"`
	Message string       `json:"message"`
}
