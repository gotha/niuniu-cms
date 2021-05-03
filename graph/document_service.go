package graph

import (
	"github.com/gotha/niuniu-cms/db"
)

type documentService interface {
	Get(id string) (*db.Document, error)
	GetAll(limit *int, offset *int, sortBy *string, sortDesc *bool) ([]db.Document, int, error)
	GetAllByTag(tagIDs []string, limit *int, offset *int, sortBy *string, sortDesc *bool) ([]db.Document, int, error)
	Create(title string, body string, tagIDs []string) (*db.Document, error)
	Update(id string, title *string, body *string, tagIDs []string) (*db.Document, error)
	Delete(id string) error
}
