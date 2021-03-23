package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/gotha/niuniu-cms/data"
	"github.com/gotha/niuniu-cms/graph/generated"
	"github.com/gotha/niuniu-cms/graph/model"
)

func (r *mutationResolver) CreateTag(ctx context.Context, input model.NewTag) (*model.Tag, error) {
	tag, err := r.tagService.New(input.Title)
	if err != nil {
		return nil, err
	}

	return TagToModel(tag), nil
}

func (r *mutationResolver) UpdateTag(ctx context.Context, id string, input model.UpdateTag) (*model.Tag, error) {
	tag, err := r.tagService.Update(id, input.Title)
	if err != nil {
		return nil, err
	}

	return TagToModel(tag), nil
}

func (r *mutationResolver) DeleteTag(ctx context.Context, id string) (bool, error) {
	err := r.tagService.Delete(id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) CreateDocument(ctx context.Context, input model.NewDocument) (*model.Document, error) {
	var tags []data.Tag
	var err error
	if input.Tags != nil {
		if len(input.Tags) > 0 {
			tags, err = r.tagService.GetMultiple(input.Tags)
			if err != nil {
				return nil, err
			}
		}
	}

	var attachments []data.Attachment
	if input.Attachments != nil {
		if len(input.Attachments) > 0 {
			attachments, err = r.attachmentService.GetMultiple(input.Attachments)
			if err != nil {
				return nil, err
			}
		}
	}

	document, err := r.documentService.New(input.Title, input.Body, tags, attachments)
	if err != nil {
		return nil, fmt.Errorf("error saving document: %w", err)
	}

	return DocumentToModel(*document), nil
}

func (r *mutationResolver) UpdateDocument(ctx context.Context, id string, input model.UpdateDocument) (*model.Document, error) {
	var tags []data.Tag
	var err error
	if input.Tags != nil {
		if len(input.Tags) > 0 {
			tags, err = r.tagService.GetMultiple(input.Tags)
			if err != nil {
				return nil, err
			}
		}
	}

	var attachments []data.Attachment
	if input.Attachments != nil {
		if len(input.Attachments) > 0 {
			attachments, err = r.attachmentService.GetMultiple(input.Attachments)
			if err != nil {
				return nil, err
			}
		}
	}

	doc, err := r.documentService.Update(id, input.Title, input.Body, tags, attachments)
	if err != nil {
		return nil, err
	}

	return DocumentToModel(*doc), nil
}

func (r *mutationResolver) DeleteDocument(ctx context.Context, id string) (bool, error) {
	err := r.documentService.Delete(id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) CreateAttachment(ctx context.Context, input model.NewAttachment) (*model.Attachment, error) {
	a, err := r.attachmentService.New(input.URL, input.Title)
	if err != nil {
		return nil, err
	}
	return AttachmentToModel(*a), nil
}

func (r *mutationResolver) UpdateAttachment(ctx context.Context, id string, input model.UpdateAttachment) (*model.Attachment, error) {
	a, err := r.attachmentService.Update(id, input.URL, input.Title)
	if err != nil {
		return nil, err
	}
	return AttachmentToModel(*a), nil
}

func (r *mutationResolver) DeleteAttachment(ctx context.Context, id string) (bool, error) {
	err := r.attachmentService.Delete(id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *queryResolver) Tags(ctx context.Context) ([]*model.Tag, error) {
	tags, err := r.tagService.GetAll()
	if err != nil {
		return nil, err
	}

	var retval []*model.Tag
	for _, tag := range tags {
		retval = append(retval, TagToModel(&tag))
	}

	return retval, nil
}

func (r *queryResolver) GetDocument(ctx context.Context, id string) (*model.Document, error) {
	doc, err := r.documentService.Get(id)
	if err != nil {
		return nil, err
	}

	return DocumentToModel(doc), nil
}

func (r *queryResolver) GetDocuments(ctx context.Context, first *int, offset *int, perPage *int, sortBy *string, sortDesc *bool) (*model.Documents, error) {
	documents, err := r.documentService.GetAll(first, offset, sortBy, sortDesc)
	if err != nil {
		return nil, err
	}

	numDocuments, err := r.documentService.GetNumDocuments()
	if err != nil {
		return nil, err
	}

	var retval []*model.Document
	for _, i := range documents {
		retval = append(retval, DocumentToModel(i))
	}

	return &model.Documents{
		Documents: retval,
		Count:     int(numDocuments),
	}, nil
}

func (r *queryResolver) GetDocumentsByTag(ctx context.Context, tagIDs []string, first *int, offset *int, sortBy *string, sortDesc *bool) (*model.Documents, error) {
	documents, err := r.documentService.GetAllByTag(tagIDs, first, offset, sortBy, sortDesc)
	if err != nil {
		return nil, err
	}

	numDocuments, err := r.documentService.GetNumDocumentsWithTag(tagIDs)
	if err != nil {
		return nil, err
	}

	var retval []*model.Document
	for _, i := range documents {
		retval = append(retval, DocumentToModel(i))
	}

	return &model.Documents{
		Documents: retval,
		Count:     int(numDocuments),
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
