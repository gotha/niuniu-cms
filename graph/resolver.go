package graph

import "github.com/gotha/niuniu-cms/data"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	tagService        *data.TagService
	documentService   *data.DocumentService
	attachmentService *data.AttachmentService
}

func NewResolver(
	tagService *data.TagService,
	documentService *data.DocumentService,
	attachmentService *data.AttachmentService,
) *Resolver {
	return &Resolver{
		tagService:        tagService,
		documentService:   documentService,
		attachmentService: attachmentService,
	}
}
