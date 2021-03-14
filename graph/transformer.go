package graph

import (
	"github.com/gotha/niuniu-cms/data"
	"github.com/gotha/niuniu-cms/graph/model"
)

func TagToModel(tag *data.Tag) *model.Tag {
	return &model.Tag{
		ID:        tag.ID.String(),
		Title:     tag.Title,
		CreatedAt: tag.CreatedAt.UTC().String(),
		UpdatedAt: tag.UpdatedAt.UTC().String(),
	}
}

func DocumentToModel(doc data.Document) *model.Document {

	var tags []*model.Tag
	for _, tag := range doc.Tags {
		tags = append(tags, TagToModel(&tag))
	}
	return &model.Document{
		ID:        doc.ID.String(),
		Title:     doc.Title,
		Body:      doc.Body,
		Tags:      tags,
		CreatedAt: doc.CreatedAt.UTC().String(),
		UpdatedAt: doc.UpdatedAt.UTC().String(),
	}
}
