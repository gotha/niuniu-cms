package data

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Title     string
	Documents []*Document `gorm:"many2many:document_tags"`
}

type TagService struct {
	db *gorm.DB
}

func NewTagService(db *gorm.DB) *TagService {
	return &TagService{
		db: db,
	}
}

func (s *TagService) GetAll() ([]Tag, error) {
	var tags []Tag
	res := s.db.Find(&tags)
	if res.Error != nil {
		return nil, res.Error
	}
	return tags, nil
}

func (s *TagService) Get(id string) (*Tag, error) {
	var tag Tag
	res := s.db.Where("id = ?", id).First(&tag)
	if res.Error != nil {
		return nil, fmt.Errorf("error getting tag with id %s: %w", id, res.Error)
	}

	return &tag, nil
}

func (s *TagService) GetMultiple(ids []string) ([]Tag, error) {
	var tags []Tag
	res := s.db.Where("id IN ?", ids).Find(&tags)
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		return nil, res.Error
	}
	return tags, nil
}

func (s *TagService) New(title string) (*Tag, error) {
	var existingTag Tag
	res := s.db.Where("title = ?", title).First(&existingTag)
	if res.RowsAffected > 0 {
		return nil, fmt.Errorf("tag with such title already exists")
	}

	tag := &Tag{
		Title: title,
	}

	res = s.db.Save(tag)
	if res.Error != nil {
		return nil, res.Error
	}
	return tag, nil
}

func (s *TagService) Update(id string, title string) (*Tag, error) {
	tag, err := s.Get(id)
	if err != nil {
		return nil, fmt.Errorf("err fetching tag %s: %w", id, err)
	}

	tag.Title = title

	res := s.db.Save(&tag)
	if res.Error != nil {
		return nil, res.Error
	}

	return tag, nil
}

func (s *TagService) Delete(id string) error {
	var tag Tag
	res := s.db.Where("id = ?", id).Delete(&tag)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
