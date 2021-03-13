package data

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
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

type DocumentService struct {
	db *gorm.DB
}

func NewDocumentService(db *gorm.DB) *DocumentService {
	return &DocumentService{
		db: db,
	}
}

func (s *DocumentService) GetAll() ([]Document, error) {
	var docs []Document
	res := s.db.Preload("Tags").Find(&docs)
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
