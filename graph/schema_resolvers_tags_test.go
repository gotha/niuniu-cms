package graph

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/gotha/niuniu-cms/data"
	"github.com/gotha/niuniu-cms/graph/model"
	"gorm.io/gorm"
)

func TestCreateTag(t *testing.T) {
	tests := []struct {
		name           string
		input          model.NewTag
		tagServiceMock *tagServiceMock
		expectedError  error
		expectedResult *model.Tag
	}{
		{
			name:  "test that error is handled",
			input: model.NewTag{Title: "mytagg"},
			tagServiceMock: &tagServiceMock{
				NewFunc: func(name string) (*data.Tag, error) {
					return nil, fmt.Errorf("service error")
				},
			},
			expectedError: fmt.Errorf("tagService was unable to create tag: service error"),
		},
		{
			name:  "test that returned model is transformed correctly",
			input: model.NewTag{Title: "mytagg"},
			tagServiceMock: &tagServiceMock{
				NewFunc: func(name string) (*data.Tag, error) {
					return &data.Tag{
						Model: gorm.Model{
							CreatedAt: time.Date(2021, time.May, 1, 10, 55, 0, 0, time.UTC),
							UpdatedAt: time.Date(2021, time.May, 1, 10, 56, 0, 0, time.UTC),
						},
						ID:    uuid.MustParse("c2ae2afe-658b-40d4-b242-a87580ec1fd7"),
						Title: "mytagg",
					}, nil
				},
			},
			expectedResult: &model.Tag{
				ID:        "c2ae2afe-658b-40d4-b242-a87580ec1fd7",
				Title:     "mytagg",
				CreatedAt: time.Date(2021, time.May, 1, 10, 55, 0, 0, time.UTC).String(),
				UpdatedAt: time.Date(2021, time.May, 1, 10, 56, 0, 0, time.UTC).String(),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resolver := NewResolver(
				test.tagServiceMock,
				&documentServiceMock{},
				&attachmentServiceMock{},
			)

			res, err := resolver.Mutation().CreateTag(context.TODO(), test.input)
			if err != nil && test.expectedError == nil {
				t.Fatalf("unexpected error received")
			}

			if test.expectedError != nil && test.expectedError.Error() != err.Error() {
				t.Fatalf("expected error '%s' got '%s'", test.expectedError.Error(), err.Error())
			}

			if res != nil && *res != *test.expectedResult {
				t.Fatalf("expected res '%v' got '%v'\n", test.expectedResult, res)
			}
		})
	}
}

func TestUpdateTag(t *testing.T) {
	tests := []struct {
		name           string
		input          model.UpdateTag
		tagServiceMock *tagServiceMock
		expectedError  error
		expectedResult *model.Tag
	}{
		{
			name:  "test that error is handled",
			input: model.UpdateTag{Title: "mytagg"},
			tagServiceMock: &tagServiceMock{
				UpdateFunc: func(id, name string) (*data.Tag, error) {
					return nil, fmt.Errorf("service error")
				},
			},
			expectedError: fmt.Errorf("tagService was unable to update tag: service error"),
		},
		{
			name:  "test that returned model is transformed correctly",
			input: model.UpdateTag{Title: "mytagg"},
			tagServiceMock: &tagServiceMock{
				UpdateFunc: func(id, name string) (*data.Tag, error) {
					return &data.Tag{
						Model: gorm.Model{
							CreatedAt: time.Date(2021, time.May, 1, 10, 55, 0, 0, time.UTC),
							UpdatedAt: time.Date(2021, time.May, 1, 10, 56, 0, 0, time.UTC),
						},
						ID:    uuid.MustParse("c2ae2afe-658b-40d4-b242-a87580ec1fd7"),
						Title: "mytagg",
					}, nil
				},
			},
			expectedResult: &model.Tag{
				ID:        "c2ae2afe-658b-40d4-b242-a87580ec1fd7",
				Title:     "mytagg",
				CreatedAt: time.Date(2021, time.May, 1, 10, 55, 0, 0, time.UTC).String(),
				UpdatedAt: time.Date(2021, time.May, 1, 10, 56, 0, 0, time.UTC).String(),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resolver := NewResolver(
				test.tagServiceMock,
				&documentServiceMock{},
				&attachmentServiceMock{},
			)

			res, err := resolver.Mutation().UpdateTag(context.TODO(), "id", test.input)
			if err != nil && test.expectedError == nil {
				t.Fatalf("unexpected error received")
			}

			if test.expectedError != nil && test.expectedError.Error() != err.Error() {
				t.Fatalf("expected error '%s' got '%s'", test.expectedError.Error(), err.Error())
			}

			if res != nil && *res != *test.expectedResult {
				t.Fatalf("expected res '%v' got '%v'\n", test.expectedResult, res)
			}
		})
	}
}

func TestDeleteTag(t *testing.T) {

	tests := []struct {
		name           string
		tagServiceMock *tagServiceMock
		expectedError  error
		expectedResult bool
	}{
		{
			name: "test that error is handled",
			tagServiceMock: &tagServiceMock{
				DeleteFunc: func(id string) error {
					return fmt.Errorf("service error")
				},
			},
			expectedError:  fmt.Errorf("tagService was unable to delete tag: service error"),
			expectedResult: false,
		},
		{
			name: "test that correct data is returned",
			tagServiceMock: &tagServiceMock{
				DeleteFunc: func(id string) error {
					return nil
				},
			},
			expectedResult: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resolver := NewResolver(
				test.tagServiceMock,
				&documentServiceMock{},
				&attachmentServiceMock{},
			)

			res, err := resolver.Mutation().DeleteTag(context.TODO(), "uuid1")
			if err != nil && test.expectedError == nil {
				t.Fatalf("unexpected error received")
			}

			if test.expectedError != nil && test.expectedError.Error() != err.Error() {
				t.Fatalf("expected error '%s' got '%s'", test.expectedError.Error(), err.Error())
			}

			if res != test.expectedResult {
				t.Fatalf("expected res '%v' got '%v'\n", test.expectedResult, res)
			}
		})
	}
}
