package document

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/gotha/niuniu-cms/db"
)

func TestGet(t *testing.T) {

	tests := []struct {
		name        string
		repoMock    *repositoryMock
		expectedRes *db.Document
		expectedErr error
	}{
		{
			name: "test that error from repo is handled",
			repoMock: &repositoryMock{
				GetFunc: func(id string) (*db.Document, error) {
					return nil, fmt.Errorf("repo error")
				},
			},
			expectedErr: fmt.Errorf("repo error"),
		},
		{
			name: "test that returning nil document is treated as error",
			repoMock: &repositoryMock{
				GetFunc: func(id string) (*db.Document, error) {
					return nil, nil
				},
			},
			expectedErr: fmt.Errorf("document not found: uuid"),
		},
		{
			name: "test that document is returned as is",
			repoMock: &repositoryMock{
				GetFunc: func(id string) (*db.Document, error) {
					return &db.Document{
						ID:    uuid.MustParse("cab498f4-0a60-413b-bba1-da4baced5fa9"),
						Title: "test1",
					}, nil
				},
			},
			expectedRes: &db.Document{
				ID:    uuid.MustParse("cab498f4-0a60-413b-bba1-da4baced5fa9"),
				Title: "test1",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := NewService(test.repoMock, nil)

			res, err := s.Get("uuid")
			if err != nil && test.expectedErr == nil {
				t.Fatalf("unexpected error: %s", err)
			}
			if test.expectedErr != nil && err.Error() != test.expectedErr.Error() {
				t.Fatalf("expected error: '%s', got '%s'", test.expectedErr, err)
			}
			if test.expectedRes != nil && !reflect.DeepEqual(*res, *test.expectedRes) {
				t.Fatalf("expected result: '%v', got '%v'", test.expectedRes, res)
			}
		})
	}
}

func TestGetAll(t *testing.T) {
	tests := []struct {
		name        string
		repoMock    *repositoryMock
		expectedRes []db.Document
		expectedNum int
		expectedErr error
	}{
		{
			name: "test that error from repo getAll is handled",
			repoMock: &repositoryMock{
				GetAllFunc: func(a, b *int, c *string, d *bool) ([]db.Document, error) {
					return nil, fmt.Errorf("repo error 1")
				},
			},
			expectedErr: fmt.Errorf("repo error 1"),
		},
		{
			name: "test that error from repo getNumDocuments is handled",
			repoMock: &repositoryMock{
				GetAllFunc: func(a, b *int, c *string, d *bool) ([]db.Document, error) {
					return nil, nil
				},
				GetNumDocumentsFunc: func() (int64, error) {
					return 0, fmt.Errorf("repo error 2")
				},
			},
			expectedErr: fmt.Errorf("repo error 2"),
		},
		{
			name: "test that everything is fine",
			repoMock: &repositoryMock{
				GetAllFunc: func(a, b *int, c *string, d *bool) ([]db.Document, error) {
					return []db.Document{
						db.Document{
							Title: "title1",
						},
					}, nil
				},
				GetNumDocumentsFunc: func() (int64, error) {
					return 64, nil
				},
			},
			expectedRes: []db.Document{
				db.Document{
					Title: "title1",
				},
			},
			expectedNum: 64,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := NewService(test.repoMock, nil)

			res, num, err := s.GetAll(nil, nil, nil, nil)
			if err != nil && test.expectedErr == nil {
				t.Fatalf("unexpected error: %s", err)
			}
			if test.expectedErr != nil && err.Error() != test.expectedErr.Error() {
				t.Fatalf("expected error: '%s', got '%s'", test.expectedErr, err)
			}
			if test.expectedRes != nil && !reflect.DeepEqual(res, test.expectedRes) {
				t.Fatalf("expected result: '%v', got '%v'", test.expectedRes, res)
			}
			if test.expectedNum != num {
				t.Fatalf("expectedNum %d, got %d", test.expectedNum, num)
			}
		})
	}
}

func TestGetAllByTag(t *testing.T) {
	tests := []struct {
		name        string
		repoMock    *repositoryMock
		expectedRes []db.Document
		expectedNum int
		expectedErr error
	}{
		{
			name: "test that error from repo getAllByTag is handled",
			repoMock: &repositoryMock{
				GetAllByTagFunc: func(tagIDs []string, a, b *int, c *string, d *bool) ([]db.Document, error) {
					return nil, fmt.Errorf("repo error 1")
				},
			},
			expectedErr: fmt.Errorf("repo error 1"),
		},
		{
			name: "test that error from repo getNumDocumentsWithTag is handled",
			repoMock: &repositoryMock{
				GetAllByTagFunc: func(tagIDs []string, a, b *int, c *string, d *bool) ([]db.Document, error) {
					return nil, nil
				},
				GetNumDocumentsWithTagFunc: func(tagIDs []string) (int64, error) {
					return 0, fmt.Errorf("repo error 2")
				},
			},
			expectedErr: fmt.Errorf("repo error 2"),
		},
		{
			name: "test that everything is fine",
			repoMock: &repositoryMock{
				GetAllByTagFunc: func(tagIDs []string, a, b *int, c *string, d *bool) ([]db.Document, error) {
					return []db.Document{
						db.Document{
							Title: "title1",
						},
					}, nil
				},
				GetNumDocumentsWithTagFunc: func(tagIDs []string) (int64, error) {
					return 64, nil
				},
			},
			expectedRes: []db.Document{
				db.Document{
					Title: "title1",
				},
			},
			expectedNum: 64,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := NewService(test.repoMock, nil)

			res, num, err := s.GetAllByTag(nil, nil, nil, nil, nil)
			if err != nil && test.expectedErr == nil {
				t.Fatalf("unexpected error: %s", err)
			}
			if test.expectedErr != nil && err.Error() != test.expectedErr.Error() {
				t.Fatalf("expected error: '%s', got '%s'", test.expectedErr, err)
			}
			if test.expectedRes != nil && !reflect.DeepEqual(res, test.expectedRes) {
				t.Fatalf("expected result: '%v', got '%v'", test.expectedRes, res)
			}
			if test.expectedNum != num {
				t.Fatalf("expectedNum %d, got %d", test.expectedNum, num)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	tests := []struct {
		name           string
		tagsIDs        []string
		repoMock       *repositoryMock
		tagServiceMock *tagServiceMock
		expectedRes    *db.Document
		expectedErr    error
	}{
		{
			name:    "test that error from tagsService is handled",
			tagsIDs: []string{"uuid1", "uuid2"},
			tagServiceMock: &tagServiceMock{
				GetMultipleFunc: func(a []string) ([]db.Tag, error) {
					return nil, fmt.Errorf("tags error")
				},
			},
			expectedErr: fmt.Errorf("tagService failed to load specified tags: tags error"),
		},
		{
			name:    "test that error is returned when tagService returns less tags than asked for",
			tagsIDs: []string{"uuid1", "uuid2"},
			tagServiceMock: &tagServiceMock{
				GetMultipleFunc: func(a []string) ([]db.Tag, error) {
					return []db.Tag{
						db.Tag{
							Title: "myTag1",
						},
					}, nil
				},
			},
			expectedErr: fmt.Errorf("some tags do not exit"),
		},
		{
			name: "test that error from saving document is handled",
			repoMock: &repositoryMock{
				CreateFunc: func(a, b string, c []db.Tag) (*db.Document, error) {
					return nil, fmt.Errorf("repo error 3")
				},
			},
			expectedErr: fmt.Errorf("error saving document: repo error 3"),
		},
		{
			name: "test that everything is fine",
			repoMock: &repositoryMock{
				CreateFunc: func(a, b string, c []db.Tag) (*db.Document, error) {
					return &db.Document{Title: "myTitle"}, nil
				},
			},
			expectedRes: &db.Document{Title: "myTitle"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := NewService(test.repoMock, test.tagServiceMock)

			res, err := s.Create("title", "body", test.tagsIDs)
			if err != nil && test.expectedErr == nil {
				t.Fatalf("unexpected error: %s", err)
			}
			if test.expectedErr != nil && err.Error() != test.expectedErr.Error() {
				t.Fatalf("expected error: '%s', got '%s'", test.expectedErr, err)
			}
			if test.expectedRes != nil && !reflect.DeepEqual(res, test.expectedRes) {
				t.Fatalf("expected result: '%v', got '%v'", test.expectedRes, res)
			}
		})
	}
}

func TestUpdate(t *testing.T) {

	pointerOfString := func(a string) *string {
		return &a
	}

	tests := []struct {
		name           string
		title          *string
		body           *string
		tagsIDs        []string
		repoMock       *repositoryMock
		tagServiceMock *tagServiceMock
		expectedRes    *db.Document
		expectedErr    error
	}{
		{
			name: "test that error is returned when document is missing",
			repoMock: &repositoryMock{
				GetFunc: func(id string) (*db.Document, error) {
					return nil, nil
				},
			},
			expectedErr: fmt.Errorf("document not found: uuid1"),
		},
		{
			name:    "test that error from tagsService is handled",
			tagsIDs: []string{"uuid1", "uuid2"},
			repoMock: &repositoryMock{
				GetFunc: func(id string) (*db.Document, error) {
					return &db.Document{}, nil
				},
			},
			tagServiceMock: &tagServiceMock{
				GetMultipleFunc: func(a []string) ([]db.Tag, error) {
					return nil, fmt.Errorf("tags error")
				},
			},
			expectedErr: fmt.Errorf("tagService failed to load specified tags: tags error"),
		},
		{
			name:    "test that error is returned when tagService returns less tags than asked for",
			tagsIDs: []string{"uuid1", "uuid2"},
			repoMock: &repositoryMock{
				GetFunc: func(id string) (*db.Document, error) {
					return &db.Document{}, nil
				},
			},
			tagServiceMock: &tagServiceMock{
				GetMultipleFunc: func(a []string) ([]db.Tag, error) {
					return []db.Tag{
						db.Tag{
							Title: "myTag1",
						},
					}, nil
				},
			},
			expectedErr: fmt.Errorf("some tags do not exit"),
		},
		{
			name: "test that error from saving document is handled",
			repoMock: &repositoryMock{
				GetFunc: func(id string) (*db.Document, error) {
					return &db.Document{}, nil
				},
				UpdateFunc: func(doc *db.Document) (*db.Document, error) {
					return nil, fmt.Errorf("repo error 4")
				},
			},
			expectedErr: fmt.Errorf("error updating document: repo error 4"),
		},
		{
			name:  "test that title is updated when provided",
			title: pointerOfString("myTitle"),
			repoMock: &repositoryMock{
				GetFunc: func(id string) (*db.Document, error) {
					return &db.Document{Title: "originalTitle"}, nil
				},
				UpdateFunc: func(doc *db.Document) (*db.Document, error) {
					return doc, nil
				},
			},
			expectedRes: &db.Document{Title: "myTitle"},
		},
		{
			name: "test that body is updated when provided",
			body: pointerOfString("newBody"),
			repoMock: &repositoryMock{
				GetFunc: func(id string) (*db.Document, error) {
					return &db.Document{
						Title: "originalTitle",
						Body:  "originalBody",
					}, nil
				},
				UpdateFunc: func(doc *db.Document) (*db.Document, error) {
					return doc, nil
				},
			},
			expectedRes: &db.Document{
				Title: "originalTitle",
				Body:  "newBody",
			},
		},
		{
			name:    "test that tags are updated when provided",
			tagsIDs: []string{"tag1", "tag2"},
			repoMock: &repositoryMock{
				GetFunc: func(id string) (*db.Document, error) {
					return &db.Document{
						Title: "originalTitle",
						Body:  "originalBody",
					}, nil
				},
				UpdateFunc: func(doc *db.Document) (*db.Document, error) {
					return doc, nil
				},
			},
			tagServiceMock: &tagServiceMock{
				GetMultipleFunc: func(a []string) ([]db.Tag, error) {
					return []db.Tag{
						{Title: "myTag1"},
						{Title: "myTag2"},
					}, nil
				},
			},
			expectedRes: &db.Document{
				Title: "originalTitle",
				Body:  "originalBody",
				Tags: []db.Tag{
					{Title: "myTag1"},
					{Title: "myTag2"},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := NewService(test.repoMock, test.tagServiceMock)

			res, err := s.Update("uuid1", test.title, test.body, test.tagsIDs)
			if err != nil && test.expectedErr == nil {
				t.Fatalf("unexpected error: %s", err)
			}
			if test.expectedErr != nil && err.Error() != test.expectedErr.Error() {
				t.Fatalf("expected error: '%s', got '%s'", test.expectedErr, err)
			}
			if test.expectedRes != nil && !reflect.DeepEqual(res, test.expectedRes) {
				t.Fatalf("expected result: '%v', got '%v'", test.expectedRes, res)
			}
		})
	}
}
