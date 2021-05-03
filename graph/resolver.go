package graph

type Resolver struct {
	tagService      tagService
	documentService documentService
}

func NewResolver(
	tagService tagService,
	documentService documentService,
) *Resolver {
	return &Resolver{
		tagService:      tagService,
		documentService: documentService,
	}
}
