package handlers

type Response struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Errors  *[]string `json:"errors"`
	Data    any       `json:"data"`
}
