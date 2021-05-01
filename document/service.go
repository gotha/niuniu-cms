package document

import (
	"fmt"

	"github.com/gotha/niuniu-cms/db"
)

const defaultLimit = 50
const defaultSortBy = "created_at"

type repository interface {
	Get(id string) (*db.Document, error)
	GetNumDocuments() (int64, error)
	GetAll(limit *int, offset *int, sortBy *string, sortDesc *bool) ([]db.Document, error)
	GetNumDocumentsWithTag(tagIDs []string) (int64, error)
	GetAllByTag(tagIDs []string, limit *int, offset *int, sortBy *string, sortDesc *bool) ([]db.Document, error)
	Create(title string, body string, tags []db.Tag) (*db.Document, error)
	Update(doc db.Document, title *string, body *string, tags []db.Tag) (*db.Document, error)
	Delete(id string) error
}

type tagService interface {
	GetMultiple(ids []string) ([]db.Tag, error)
}

type Service struct {
	repo       repository
	tagService tagService
}

func NewService(repo repository, tagService tagService) *Service {
	return &Service{
		repo:       repo,
		tagService: tagService,
	}
}

func (s *Service) Get(id string) (*db.Document, error) {
	doc, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	if doc == nil {
		return nil, fmt.Errorf("document not found: %s", id)
	}

	return doc, nil
}

func (s *Service) GetAll(limit *int, offset *int, sortBy *string, sortDesc *bool) ([]db.Document, int, error) {
	docs, err := s.repo.GetAll(limit, offset, sortBy, sortDesc)
	if err != nil {
		return nil, 0, err
	}

	numDocuments, err := s.repo.GetNumDocuments()
	if err != nil {
		return nil, 0, err
	}

	return docs, int(numDocuments), nil
}

func (s *Service) GetAllByTag(tagIDs []string, limit *int, offset *int, sortBy *string, sortDesc *bool) ([]db.Document, int, error) {

	docs, err := s.repo.GetAllByTag(tagIDs, limit, offset, sortBy, sortDesc)
	if err != nil {
		return nil, 0, err
	}

	numDocuments, err := s.repo.GetNumDocumentsWithTag(tagIDs)
	if err != nil {
		return nil, 0, err
	}

	return docs, int(numDocuments), nil
}

func (s *Service) Create(title string, body string, tagIDs []string) (*db.Document, error) {
	var tags []db.Tag
	var err error
	if tagIDs != nil {
		if len(tagIDs) > 0 {
			tags, err = s.tagService.GetMultiple(tagIDs)
			if err != nil {
				return nil, fmt.Errorf("tagService failed to load specified tags: %w", err)
			}
		}
	}

	doc, err := s.repo.Create(title, body, tags)
	if err != nil {
		return nil, fmt.Errorf("error saving document: %w", err)
	}
	return doc, nil
}

func (s *Service) Update(id string, title *string, body *string, tagIDs []string) (*db.Document, error) {
	doc, err := s.Get(id)
	if err != nil {
		return nil, err
	}

	var tags []db.Tag
	if tagIDs != nil {
		if len(tagIDs) > 0 {
			tags, err = s.tagService.GetMultiple(tagIDs)
			if err != nil {
				return nil, fmt.Errorf("tagService failed to load specified tags: %w", err)
			}
		}
	}

	// @todo - this looks bad
	res, err := s.repo.Update(*doc, title, body, tags)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Service) Delete(id string) error {
	return s.repo.Delete(id)
}
