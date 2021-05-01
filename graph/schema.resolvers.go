package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/gotha/niuniu-cms/graph/generated"
	"github.com/gotha/niuniu-cms/graph/model"
)

func (r *mutationResolver) CreateTag(ctx context.Context, input model.NewTag) (*model.Tag, error) {
	tag, err := r.tagService.Create(input.Title)
	if err != nil {
		return nil, fmt.Errorf("tagService was unable to create tag: %w", err)
	}

	return TagToModel(tag), nil
}

func (r *mutationResolver) UpdateTag(ctx context.Context, id string, input model.UpdateTag) (*model.Tag, error) {
	tag, err := r.tagService.Update(id, input.Title)
	if err != nil {
		return nil, fmt.Errorf("tagService was unable to update tag: %w", err)
	}

	return TagToModel(tag), nil
}

func (r *mutationResolver) DeleteTag(ctx context.Context, id string) (bool, error) {
	err := r.tagService.Delete(id)
	if err != nil {
		return false, fmt.Errorf("tagService was unable to delete tag: %w", err)
	}
	return true, nil
}

func (r *mutationResolver) CreateDocument(ctx context.Context, input model.NewDocument) (*model.Document, error) {
	doc, err := r.documentService.Create(input.Title, input.Body, input.Tags)
	if err != nil {
		return nil, err
	}

	return DocumentToModel(*doc), nil
}

func (r *mutationResolver) UpdateDocument(ctx context.Context, id string, input model.UpdateDocument) (*model.Document, error) {
	doc, err := r.documentService.Update(id, input.Title, input.Body, input.Tags)
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

func (r *queryResolver) Tags(ctx context.Context) ([]*model.Tag, error) {
	tags, err := r.tagService.GetAll()
	if err != nil {
		return nil, err
	}

	var retval []*model.Tag
	for i := range tags {
		retval = append(retval, TagToModel(&tags[i]))
	}

	return retval, nil
}

func (r *queryResolver) GetDocument(ctx context.Context, id string) (*model.Document, error) {
	doc, err := r.documentService.Get(id)
	if err != nil {
		return nil, err
	}

	return DocumentToModel(*doc), nil
}

func (r *queryResolver) GetDocuments(ctx context.Context, first *int, offset *int, perPage *int, sortBy *string, sortDesc *bool) (*model.Documents, error) {
	documents, num, err := r.documentService.GetAll(first, offset, sortBy, sortDesc)
	if err != nil {
		return nil, err
	}

	var retval []*model.Document
	for _, i := range documents {
		retval = append(retval, DocumentToModel(i))
	}

	return &model.Documents{
		Documents: retval,
		Count:     num,
	}, nil
}

func (r *queryResolver) GetDocumentsByTag(ctx context.Context, tagIDs []string, first *int, offset *int, sortBy *string, sortDesc *bool) (*model.Documents, error) {
	documents, num, err := r.documentService.GetAllByTag(tagIDs, first, offset, sortBy, sortDesc)
	if err != nil {
		return nil, err
	}

	var retval []*model.Document
	for _, i := range documents {
		retval = append(retval, DocumentToModel(i))
	}

	return &model.Documents{
		Documents: retval,
		Count:     num,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
