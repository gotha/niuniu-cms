package document

//go:generate moq -out ./tag_service_mock.go . tagService

import "github.com/gotha/niuniu-cms/db"

type tagService interface {
	GetMultiple(ids []string) ([]db.Tag, error)
}
