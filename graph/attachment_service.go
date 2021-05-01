package graph

//go:generate moq -out ./attachment_service_mock.go . attachmentService

import "github.com/gotha/niuniu-cms/data"

type attachmentService interface {
	Get(id string) (*data.Attachment, error)
	GetMultiple(ids []string) ([]data.Attachment, error)
	New(url string, title *string) (*data.Attachment, error)
	Update(id, url string, title *string) (*data.Attachment, error)
	Delete(id string) error
}
