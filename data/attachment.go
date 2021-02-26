package data

import (
	"time"

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
