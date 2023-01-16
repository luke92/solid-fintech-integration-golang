package models

type List[T any] struct {
	Total int `json:"total"`
	Data  []T `json:"data"`
}
