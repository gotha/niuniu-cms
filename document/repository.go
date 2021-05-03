package document

//go:generate moq -out ./repository_mock.go . repository

import (
	"github.com/gotha/niuniu-cms/db"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type repository interface {
	Get(id string) (*db.Document, error)
	GetNumDocuments() (int64, error)
	GetAll(limit *int, offset *int, sortBy *string, sortDesc *bool) ([]db.Document, error)
	GetNumDocumentsWithTag(tagIDs []string) (int64, error)
	GetAllByTag(tagIDs []string, limit *int, offset *int, sortBy *string, sortDesc *bool) ([]db.Document, error)
	Create(title string, body string, tags []db.Tag) (*db.Document, error)
	Update(doc *db.Document) (*db.Document, error)
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

func (r *Repository) Get(id string) (*db.Document, error) {
	var doc db.Document
	res := r.db.
		Preload("Tags").
		Where("documents.id = ?", id).
		First(&doc)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, res.Error
	}

	return &doc, nil
}

func (r *Repository) GetNumDocuments() (int64, error) {
	var num int64
	res := r.db.Model(&db.Document{}).Count(&num)
	if res.Error != nil {
		return 0, res.Error
	}
	return num, nil
}

func (r *Repository) GetAll(limit *int, offset *int, sortBy *string, sortDesc *bool) ([]db.Document, error) {
	query := r.db.Preload("Tags")

	limitDocuments := defaultLimit
	if limit != nil {
		limitDocuments = *limit
	}
	query = query.Limit(limitDocuments)

	if offset != nil {
		query = query.Offset(*offset)
	}

	sortColumn := defaultSortBy
	if sortBy != nil {
		sortColumn = *sortBy
	}
	sortDescB := true
	if sortDesc != nil && !*sortDesc {
		sortDescB = false
	}

	query = query.Order(clause.OrderByColumn{
		Column: clause.Column{Name: sortColumn},
		Desc:   sortDescB,
	})

	var docs []db.Document
	res := query.Find(&docs)
	if res.Error != nil {
		return nil, res.Error
	}

	return docs, nil
}

func (r *Repository) GetNumDocumentsWithTag(tagIDs []string) (int64, error) {
	var num int64
	res := r.db.Model(&db.Document{}).
		Select("COUNT(DISTINCT dt.document_id)").
		Joins("JOIN document_tags AS dt ON dt.document_id = documents.id").
		Where("dt.tag_id IN ?", tagIDs).
		Count(&num)
	if res.Error != nil {
		return 0, res.Error
	}
	return num, nil
}

func (r *Repository) GetAllByTag(tagIDs []string, limit *int, offset *int, sortBy *string, sortDesc *bool) ([]db.Document, error) {
	query := r.db.Preload("Tags")

	limitDocuments := defaultLimit
	if limit != nil {
		limitDocuments = *limit
	}
	query = query.Limit(limitDocuments)

	if offset != nil {
		query = query.Offset(*offset)
	}

	sortColumn := defaultSortBy
	if sortBy != nil {
		sortColumn = *sortBy
	}
	sortDescB := true
	if sortDesc != nil && !*sortDesc {
		sortDescB = false
	}

	query = query.Order(clause.OrderByColumn{
		Column: clause.Column{Name: sortColumn},
		Desc:   sortDescB,
	})

	var docs []db.Document
	res := query.Joins("JOIN document_tags AS dt ON dt.document_id = documents.id").
		Where("dt.tag_id IN ?", tagIDs).
		Find(&docs)
	if res.Error != nil {
		return nil, res.Error
	}

	return docs, nil
}

func (r *Repository) Create(title string, body string, tags []db.Tag) (*db.Document, error) {
	doc := &db.Document{
		Title: title,
		Body:  body,
	}
	doc.Tags = append(doc.Tags, tags...)

	res := r.db.Save(doc)
	if res.Error != nil {
		return nil, res.Error
	}
	return doc, nil
}

func (r *Repository) Update(doc *db.Document) (*db.Document, error) {
	res := r.db.Save(&doc)
	if res.Error != nil {
		return nil, res.Error
	}

	return doc, nil
}

func (r *Repository) Delete(id string) error {
	res := r.db.Where("id = ?", id).Delete(&db.Document{})
	if res.Error != nil {
		return res.Error
	}
	return nil
}
