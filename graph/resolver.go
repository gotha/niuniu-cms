package graph

import "github.com/gotha/niuniu-cms/data"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	tagService      *data.TagService
	documentService *data.DocumentService
}

func NewResolver(
	tagService *data.TagService,
	documentService *data.DocumentService,
) *Resolver {
	return &Resolver{
		tagService:      tagService,
		documentService: documentService,
	}
}
