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
	tag, err := r.tagService.New(input.Title)
	if err != nil {
		return nil, err
	}

	return &model.Tag{
		ID:    tag.ID.String(),
		Title: input.Title,
	}, nil
}

func (r *mutationResolver) UpdateTag(ctx context.Context, id string, input model.UpdateTag) (*model.Tag, error) {

	tag, err := r.tagService.Update(id, input.Title)
	if err != nil {
		return nil, err
	}

	return &model.Tag{
		ID:    tag.ID.String(),
		Title: tag.Title,
	}, nil
}

func (r *mutationResolver) DeleteTag(ctx context.Context, id string) (bool, error) {

	err := r.tagService.Delete(id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) CreateDocument(ctx context.Context, input model.NewDocument) (*model.Document, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateDocument(ctx context.Context, id string, input model.UpdateDocument) (*model.Document, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteDocument(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Tags(ctx context.Context) ([]*model.Tag, error) {
	tags, err := r.tagService.GetAll()
	if err != nil {
		return nil, err
	}

	var retval []*model.Tag
	for _, i := range tags {
		retval = append(retval, &model.Tag{
			ID:    i.ID.String(),
			Title: i.Title,
		})
	}

	return retval, nil
}

func (r *queryResolver) Documents(ctx context.Context) ([]*model.Document, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
