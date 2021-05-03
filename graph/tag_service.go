package graph

import "github.com/gotha/niuniu-cms/db"

type tagService interface {
	GetAll() ([]db.Tag, error)
	Get(id string) (*db.Tag, error)
	GetMultiple(ids []string) ([]db.Tag, error)
	Create(title string) (*db.Tag, error)
	Update(id string, title string) (*db.Tag, error)
	Delete(id string) error
}
