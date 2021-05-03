package tag

//go:generate moq -out ./repository_mock.go . repository

import (
	"fmt"

	"github.com/gotha/niuniu-cms/db"
	"gorm.io/gorm"
)

type repository interface {
	GetAll() ([]db.Tag, error)
	Get(id string) (*db.Tag, error)
	GetTagByTitle(title string) (*db.Tag, error)
	GetMultiple(ids []string) ([]db.Tag, error)
	Create(title string) (*db.Tag, error)
	Update(tag db.Tag, title string) (*db.Tag, error)
	Delete(id string) error
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetAll() ([]db.Tag, error) {
	var tags []db.Tag
	res := r.db.Find(&tags)
	if res.Error != nil {
		return nil, res.Error
	}
	return tags, nil
}

func (r *Repository) Get(id string) (*db.Tag, error) {
	var tag db.Tag
	res := r.db.Where("id = ?", id).First(&tag)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("error getting tag with id %s: %w", id, res.Error)
	}

	return &tag, nil
}

func (r *Repository) GetTagByTitle(title string) (*db.Tag, error) {
	var tag db.Tag
	res := r.db.Where("title = ?", title).First(&tag)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("error getting tag with title %s: %w", title, res.Error)
	}
	return &tag, nil
}

func (r *Repository) GetMultiple(ids []string) ([]db.Tag, error) {
	var tags []db.Tag
	res := r.db.Where("id IN ?", ids).Find(&tags)
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		return nil, res.Error
	}
	return tags, nil
}

func (r *Repository) Create(title string) (*db.Tag, error) {
	tag := &db.Tag{
		Title: title,
	}
	res := r.db.Save(tag)
	if res.Error != nil {
		return nil, res.Error
	}
	return tag, nil
}

func (r *Repository) Update(tag db.Tag, title string) (*db.Tag, error) {
	tag.Title = title

	res := r.db.Save(&tag)
	if res.Error != nil {
		return nil, res.Error
	}

	return &tag, nil
}

func (r *Repository) Delete(id string) error {
	var tag db.Tag
	res := r.db.Where("id = ?", id).Delete(&tag)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
