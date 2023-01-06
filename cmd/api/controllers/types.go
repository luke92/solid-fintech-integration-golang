package controllers

type ErrorResponse struct {
	Title   string `json:"error"`
	Message string `json:"msg"`
}
