package data

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Document struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Title       string
	Body        string
	CreatedAt   time.Time    `gorm:"autoCreateTime"`
	UpdatedAt   time.Time    `gorm:"autoUpdateTime"`
	Tags        []Tag        `gorm:"many2many:document_tags;"`
	Attachments []Attachment `gorm:"many2many:document_attachments;"`
}

const defaultLimit = 50
const defaultSortBy = "created_at"

type DocumentService struct {
	db *gorm.DB
}

func NewDocumentService(db *gorm.DB) *DocumentService {
	return &DocumentService{
		db: db,
	}
}

func (s *DocumentService) Get(UUID string) (Document, error) {

	var doc Document
	res := s.db.
		Preload("Tags").
		Where("documents.id = ?", UUID).
		First(&doc)
	if res.Error != nil {
		return Document{}, res.Error
	}

	return doc, nil
}

func (s *DocumentService) GetNumDocuments() (int64, error) {
	var num int64
	res := s.db.Model(&Document{}).Count(&num)
	if res.Error != nil {
		return 0, res.Error
	}
	return num, nil
}

func (s *DocumentService) GetAll(limit *int, offset *int, sortBy *string, sortDesc *bool) ([]Document, error) {

	query := s.db.Preload("Tags")

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
	if sortDesc != nil && *sortDesc == false {
		sortDescB = false
	}

	query = query.Order(clause.OrderByColumn{
		Column: clause.Column{Name: sortColumn},
		Desc:   sortDescB,
	})

	var docs []Document
	res := query.Find(&docs)
	if res.Error != nil {
		return nil, res.Error
	}

	return docs, nil
}

func (s *DocumentService) GetNumDocumentsWithTag(tagIDs []string) (int64, error) {
	var num int64
	res := s.db.Model(&Document{}).
		Joins("JOIN document_tags AS dt ON dt.document_id = documents.id").
		Where("dt.tag_id IN ?", tagIDs).
		Count(&num)
	if res.Error != nil {
		return 0, res.Error
	}
	return num, nil
}

func (s *DocumentService) GetAllByTag(tagIDs []string, limit *int, offset *int, sortBy *string, sortDesc *bool) ([]Document, error) {

	query := s.db.Preload("Tags")

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
	if sortDesc != nil && *sortDesc == false {
		sortDescB = false
	}

	query = query.Order(clause.OrderByColumn{
		Column: clause.Column{Name: sortColumn},
		Desc:   sortDescB,
	})

	var docs []Document
	res := query.Joins("JOIN document_tags AS dt ON dt.document_id = documents.id").
		Where("dt.tag_id IN ?", tagIDs).
		Find(&docs)
	if res.Error != nil {
		return nil, res.Error
	}

	return docs, nil
}

func (s *DocumentService) New(title string, body string, tags []Tag) (*Document, error) {
	doc := &Document{
		Title: title,
		Body:  body,
	}
	doc.Tags = append(doc.Tags, tags...)

	res := s.db.Save(doc)
	if res.Error != nil {
		return nil, res.Error
	}
	return doc, nil
}

func (s *DocumentService) Update(id string, title string) (*Document, error) {
	panic("not implemented")
}

func (s *DocumentService) Delete(id string) error {
	panic("not implemented")
}
