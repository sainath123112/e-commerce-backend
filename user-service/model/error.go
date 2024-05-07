package model

type ErrorResponse struct {
	Message     string `json:"message"`
	ErrorString string `json:"error"`
}
