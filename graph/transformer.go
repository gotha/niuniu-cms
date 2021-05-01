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
	for i := range doc.Tags {
		tags = append(tags, TagToModel(&doc.Tags[i]))
	}

	var attachments []*model.Attachment
	for _, a := range doc.Attachments {
		attachments = append(attachments, AttachmentToModel(a))
	}
	return &model.Document{
		ID:          doc.ID.String(),
		Title:       doc.Title,
		Body:        doc.Body,
		Tags:        tags,
		Attachments: attachments,
		CreatedAt:   doc.CreatedAt.UTC().String(),
		UpdatedAt:   doc.UpdatedAt.UTC().String(),
	}
}

func AttachmentToModel(a data.Attachment) *model.Attachment {
	return &model.Attachment{
		ID:        a.ID.String(),
		Title:     &a.Title,
		URL:       a.URL,
		CreatedAt: a.CreatedAt.UTC().String(),
		UpdatedAt: a.UpdatedAt.UTC().String(),
	}
}
