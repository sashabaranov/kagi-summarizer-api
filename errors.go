package kagi

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
