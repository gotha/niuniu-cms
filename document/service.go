package document

import (
	"fmt"

	"github.com/gotha/niuniu-cms/db"
)

const defaultLimit = 50
const defaultSortBy = "created_at"

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
			if len(tags) < len(tagIDs) {
				return nil, fmt.Errorf("some tags do not exit")
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

	if title != nil {
		doc.Title = *title
	}
	if body != nil {
		doc.Body = *body
	}

	var tags []db.Tag
	if tagIDs != nil {
		if len(tagIDs) > 0 {
			tags, err = s.tagService.GetMultiple(tagIDs)
			if err != nil {
				return nil, fmt.Errorf("tagService failed to load specified tags: %w", err)
			}
			if len(tags) < len(tagIDs) {
				return nil, fmt.Errorf("some tags do not exit")
			}
			doc.Tags = tags
		}
	}

	res, err := s.repo.Update(doc)
	if err != nil {
		return nil, fmt.Errorf("error updating document: %w", err)
	}

	return res, nil
}

func (s *Service) Delete(id string) error {
	return s.repo.Delete(id)
}
