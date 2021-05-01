package graph

//go:generate moq -out ./document_service_mock.go . documentService

import "github.com/gotha/niuniu-cms/data"

type documentService interface {
	Get(id string) (data.Document, error)
	GetNumDocuments() (int64, error)
	GetAll(limit *int, offset *int, sortBy *string, sortDesc *bool) ([]data.Document, error)
	GetNumDocumentsWithTag(tagIDs []string) (int64, error)
	GetAllByTag(tagIDs []string, limit *int, offset *int, sortBy *string, sortDesc *bool) ([]data.Document, error)
	New(title string, body string, tags []data.Tag, attachments []data.Attachment) (*data.Document, error)
	Update(id string, title *string, body *string, tags []data.Tag, attachments []data.Attachment) (*data.Document, error)
	Delete(id string) error
}
