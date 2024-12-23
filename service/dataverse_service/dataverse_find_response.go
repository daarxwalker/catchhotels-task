package dataverse_service

type FindResponse[T any] struct {
	Value []T `json:"value"`
}
