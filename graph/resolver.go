package graph

type Resolver struct {
	tagService        tagService
	documentService   documentService
	attachmentService attachmentService
}

func NewResolver(
	tagService tagService,
	documentService documentService,
	attachmentService attachmentService,
) *Resolver {
	return &Resolver{
		tagService:        tagService,
		documentService:   documentService,
		attachmentService: attachmentService,
	}
}
