package graph

//go:generate moq -out ./tag_service_mock.go . tagService

import "github.com/gotha/niuniu-cms/data"

type tagService interface {
	GetAll() ([]data.Tag, error)
	Get(id string) (*data.Tag, error)
	GetMultiple(ids []string) ([]data.Tag, error)
	New(title string) (*data.Tag, error)
	Update(id string, title string) (*data.Tag, error)
	Delete(id string) error
}
