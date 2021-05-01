package tag

import (
	"fmt"

	"github.com/gotha/niuniu-cms/db"
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

type Service struct {
	repo repository
}

func NewService(repo repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetAll() ([]db.Tag, error) {
	return s.repo.GetAll()
}

func (s *Service) Get(id string) (*db.Tag, error) {
	tag, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	if tag == nil {
		return nil, fmt.Errorf("tag not found: %s", id)
	}
	return tag, nil
}

func (s *Service) GetMultiple(ids []string) ([]db.Tag, error) {
	return s.repo.GetMultiple(ids)
}

func (s *Service) Create(title string) (*db.Tag, error) {
	tag, err := s.repo.GetTagByTitle(title)
	if err != nil {
		return nil, fmt.Errorf("error checking if tag already exists: %w", err)
	}
	if tag != nil {
		return nil, fmt.Errorf("cannot create tag with title %s, already exists", title)
	}

	return s.repo.Create(title)
}

func (s *Service) Update(id string, title string) (*db.Tag, error) {
	tag, err := s.repo.Get(id)
	if err != nil {
		return nil, fmt.Errorf("err fetching tag %s: %w", id, err)
	}
	if tag == nil {
		return nil, fmt.Errorf("trying to update non existing tag: %s", id)
	}

	res, err := s.repo.Update(*tag, title)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Service) Delete(id string) error {
	return s.repo.Delete(id)
}
