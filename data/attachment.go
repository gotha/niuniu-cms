package data

import (
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Attachment struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Title     string
	URL       string
	CreatedAt time.Time   `gorm:"autoCreateTime"`
	UpdatedAt time.Time   `gorm:"autoUpdateTime"`
	Documents []*Document `gorm:"many2many:document_attachments"`
}

type AttachmentService struct {
	db *gorm.DB
}

func NewAttachmentService(db *gorm.DB) *AttachmentService {
	return &AttachmentService{
		db: db,
	}
}

func (s *AttachmentService) Get(ID string) (*Attachment, error) {
	var a *Attachment
	res := s.db.Where("id = ?", ID).First(a)
	if res.RowsAffected < 1 {
		return nil, fmt.Errorf("no such attachment")
	}
	if res.Error != nil {
		return nil, res.Error
	}
	return a, nil
}

func (s *AttachmentService) GetMultiple(IDs []string) ([]Attachment, error) {
	var attachments []Attachment
	res := s.db.Where("id IN ?", IDs).Find(&attachments)
	if res.Error != nil {
		return nil, res.Error
	}
	return attachments, nil
}

func (s *AttachmentService) New(url string, title *string) (*Attachment, error) {
	var existingAttachment Attachment
	res := s.db.Where("url = ?", url).First(&existingAttachment)
	if res.RowsAffected > 0 {
		return nil, fmt.Errorf("attachment with such url already exists")
	}
	spew.Dump(res.Error)
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		return nil, res.Error
	}

	a := &Attachment{URL: url}
	if title != nil {
		a.Title = *title
	}

	res = s.db.Save(a)
	if res.Error != nil {
		return nil, res.Error
	}
	return a, nil
}

func (s *AttachmentService) Update(id, url string, title *string) (*Attachment, error) {
	var a Attachment
	res := s.db.Where("id = ?", id).First(&a)
	if res.RowsAffected < 1 {
		return nil, fmt.Errorf("no such attachment")
	}
	if res.Error != nil {
		return nil, res.Error
	}

	a.URL = url
	if title != nil {
		a.Title = *title
	}

	res = s.db.Save(a)
	if res.Error != nil {
		return nil, res.Error
	}
	return &a, nil
}

func (s *AttachmentService) Delete(id string) error {
	var a Attachment
	res := s.db.Where("id = ?", id).Delete(&a)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
