package model

type Response[T any] struct {
	Success bool `json:"success"`
	Result  T    `json:"result"`
	// ErrorMessage
	Error string `json:"error"`
}
