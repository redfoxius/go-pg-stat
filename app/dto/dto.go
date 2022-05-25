package dto

import "github.com/redfoxius/go-pg-stat/app/repositories/stat/models"

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Response struct {
	Result *models.StatResult `json:"result"`
}
