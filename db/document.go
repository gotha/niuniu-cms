package db

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Document struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Title     string
	Body      string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Tags      []Tag     `gorm:"many2many:document_tags;"`
}

func (a *Document) Equal(b *Document) bool {
	if a.ID.String() == b.ID.String() &&
		a.Title == b.Title &&
		a.Body == b.Body &&
		a.CreatedAt.String() == b.CreatedAt.String() &&
		a.UpdatedAt.String() == b.UpdatedAt.String() &&
		tagsEqual(a.Tags, b.Tags) {
		return true
	}
	return false
}

func tagsEqual(a, b []Tag) bool {
	if len(a) != len(b) {
		return false
	}
	found := 0
	for i := range a {
		for y := range b {
			if a[i].Equal(&b[y]) {
				found++
				break
			}
		}
	}

	return len(a) == found
}
