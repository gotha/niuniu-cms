package tag

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/gotha/niuniu-cms/db"
)

func TestServiceGetAll(t *testing.T) {
	tests := []struct {
		name        string
		repoMock    *repositoryMock
		expectedErr error
	}{
		{
			name: "test that error is propagated",
			repoMock: &repositoryMock{
				GetAllFunc: func() ([]db.Tag, error) {
					return nil, fmt.Errorf("errX")
				},
			},
			expectedErr: fmt.Errorf("errX"),
		},
		{
			name: "test that it works",
			repoMock: &repositoryMock{
				GetAllFunc: func() ([]db.Tag, error) {
					return []db.Tag{}, nil
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := NewService(test.repoMock)
			_, err := s.GetAll()
			if err != nil {
				if test.expectedErr == nil {
					t.Fatalf("unexpected error: %s", err)
				}
				if err.Error() != test.expectedErr.Error() {
					t.Fatalf("expected err: '%s', got '%s'", test.expectedErr, err)
				}
			}
		})
	}
}

func TestServiceGetMultiple(t *testing.T) {
	tests := []struct {
		name        string
		repoMock    *repositoryMock
		expectedErr error
	}{
		{
			name: "test that error is propagated",
			repoMock: &repositoryMock{
				GetMultipleFunc: func(x []string) ([]db.Tag, error) {
					return nil, fmt.Errorf("errX")
				},
			},
			expectedErr: fmt.Errorf("errX"),
		},
		{
			name: "test that it works",
			repoMock: &repositoryMock{
				GetMultipleFunc: func(x []string) ([]db.Tag, error) {
					return []db.Tag{}, nil
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := NewService(test.repoMock)
			_, err := s.GetMultiple(nil)
			if err != nil {
				if test.expectedErr == nil {
					t.Fatalf("unexpected error: %s", err)
				}
				if err.Error() != test.expectedErr.Error() {
					t.Fatalf("expected err: '%s', got '%s'", test.expectedErr, err)
				}
			}
		})
	}
}

func TestServiceGet(t *testing.T) {
	tests := []struct {
		name        string
		repoMock    *repositoryMock
		expectedRes *db.Tag
		expectedErr error
	}{
		{
			name: "test that error from repo is handled",
			repoMock: &repositoryMock{
				GetFunc: func(a string) (*db.Tag, error) {
					return nil, fmt.Errorf("repo err 1")
				},
			},
			expectedErr: fmt.Errorf("repo err 1"),
		},
		{
			name: "test that error is returned when tag is non existent",
			repoMock: &repositoryMock{
				GetFunc: func(a string) (*db.Tag, error) {
					return nil, nil
				},
			},
			expectedErr: fmt.Errorf("tag not found: id1"),
		},
		{
			name: "test that it works",
			repoMock: &repositoryMock{
				GetFunc: func(a string) (*db.Tag, error) {
					return &db.Tag{
						ID:    uuid.MustParse("4a94b51f-a090-4625-aaa1-2bef5c680d2f"),
						Title: "testTag1",
					}, nil
				},
			},
			expectedRes: &db.Tag{
				ID:    uuid.MustParse("4a94b51f-a090-4625-aaa1-2bef5c680d2f"),
				Title: "testTag1",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := NewService(test.repoMock)
			res, err := s.Get("id1")
			if err != nil {
				if test.expectedErr == nil {
					t.Fatalf("unexpected error: %s", err)
				}
				if err.Error() != test.expectedErr.Error() {
					t.Fatalf("expected err: '%s', got '%s'", test.expectedErr, err)
				}
			}

			if res != nil && !res.Equal(test.expectedRes) {
				t.Fatalf("expected tag: '%v', got '%v'", test.expectedErr, res)
			}
		})
	}
}

func TestServiceCreate(t *testing.T) {
	tests := []struct {
		name        string
		repoMock    *repositoryMock
		expectedRes *db.Tag
		expectedErr error
	}{
		{
			name: "test that error from repo is handled",
			repoMock: &repositoryMock{
				GetTagByTitleFunc: func(a string) (*db.Tag, error) {
					return nil, fmt.Errorf("repo err 1")
				},
			},
			expectedErr: fmt.Errorf("error checking if tag already exists: repo err 1"),
		},
		{
			name: "test that error is returned when tag exists",
			repoMock: &repositoryMock{
				GetTagByTitleFunc: func(a string) (*db.Tag, error) {
					return &db.Tag{}, nil
				},
			},
			expectedErr: fmt.Errorf("cannot create tag with title tag1, already exists"),
		},
		{
			name: "test that error from create is propagated",
			repoMock: &repositoryMock{
				GetTagByTitleFunc: func(a string) (*db.Tag, error) {
					return nil, nil
				},
				CreateFunc: func(a string) (*db.Tag, error) {
					return nil, fmt.Errorf("repo error 2")
				},
			},
			expectedErr: fmt.Errorf("repo error 2"),
		},
		{
			name: "test that it works",
			repoMock: &repositoryMock{
				GetTagByTitleFunc: func(a string) (*db.Tag, error) {
					return nil, nil
				},
				CreateFunc: func(a string) (*db.Tag, error) {
					return &db.Tag{
						ID:    uuid.MustParse("4a94b51f-a090-4625-aaa1-2bef5c680d2f"),
						Title: "testTag1",
					}, nil
				},
			},
			expectedRes: &db.Tag{
				ID:    uuid.MustParse("4a94b51f-a090-4625-aaa1-2bef5c680d2f"),
				Title: "testTag1",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := NewService(test.repoMock)
			res, err := s.Create("tag1")
			if err != nil {
				if test.expectedErr == nil {
					t.Fatalf("unexpected error: %s", err)
				}
				if err.Error() != test.expectedErr.Error() {
					t.Fatalf("expected err: '%s', got '%s'", test.expectedErr, err)
				}
			}

			if res != nil && !res.Equal(test.expectedRes) {
				t.Fatalf("expected tag: '%v', got '%v'", test.expectedErr, res)
			}
		})
	}
}

func TestServiceUpdate(t *testing.T) {
	tests := []struct {
		name        string
		repoMock    *repositoryMock
		expectedRes *db.Tag
		expectedErr error
	}{
		{
			name: "test that error from repo is handled",
			repoMock: &repositoryMock{
				GetFunc: func(a string) (*db.Tag, error) {
					return nil, fmt.Errorf("repo err 1")
				},
			},
			expectedErr: fmt.Errorf("err fetching tag id1: repo err 1"),
		},
		{
			name: "test that error is returned when tag does not exist",
			repoMock: &repositoryMock{
				GetFunc: func(a string) (*db.Tag, error) {
					return nil, nil
				},
			},
			expectedErr: fmt.Errorf("trying to update non existing tag: id1"),
		},
		{
			name: "test that error from repo update is propagated",
			repoMock: &repositoryMock{
				GetFunc: func(a string) (*db.Tag, error) {
					return &db.Tag{}, nil
				},
				UpdateFunc: func(a db.Tag, b string) (*db.Tag, error) {
					return nil, fmt.Errorf("repo error 2")
				},
			},
			expectedErr: fmt.Errorf("repo error 2"),
		},
		{
			name: "test that it works",
			repoMock: &repositoryMock{
				GetFunc: func(a string) (*db.Tag, error) {
					return &db.Tag{}, nil
				},
				UpdateFunc: func(t db.Tag, a string) (*db.Tag, error) {
					return &db.Tag{
						ID:    uuid.MustParse("4a94b51f-a090-4625-aaa1-2bef5c680d2f"),
						Title: "testTag1",
					}, nil
				},
			},
			expectedRes: &db.Tag{
				ID:    uuid.MustParse("4a94b51f-a090-4625-aaa1-2bef5c680d2f"),
				Title: "testTag1",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := NewService(test.repoMock)
			res, err := s.Update("id1", "tag1")
			if err != nil {
				if test.expectedErr == nil {
					t.Fatalf("unexpected error: %s", err)
				}
				if err.Error() != test.expectedErr.Error() {
					t.Fatalf("expected err: '%s', got '%s'", test.expectedErr, err)
				}
			}

			if res != nil && !res.Equal(test.expectedRes) {
				t.Fatalf("expected tag: '%v', got '%v'", test.expectedErr, res)
			}
		})
	}
}

func TestServiceDelete(t *testing.T) {
	tests := []struct {
		name        string
		repoMock    *repositoryMock
		expectedErr error
	}{
		{
			name: "test that error is propagated",
			repoMock: &repositoryMock{
				DeleteFunc: func(x string) error {
					return fmt.Errorf("errX")
				},
			},
			expectedErr: fmt.Errorf("errX"),
		},
		{
			name: "test that it works",
			repoMock: &repositoryMock{
				DeleteFunc: func(x string) error {
					return nil
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := NewService(test.repoMock)
			err := s.Delete("")
			if err != nil {
				if test.expectedErr == nil {
					t.Fatalf("unexpected error: %s", err)
				}
				if err.Error() != test.expectedErr.Error() {
					t.Fatalf("expected err: '%s', got '%s'", test.expectedErr, err)
				}
			}
		})
	}
}
