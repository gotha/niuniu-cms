// +build integration

package document

import (
	"fmt"
	"reflect"

	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/gotha/niuniu-cms/db"
	"gorm.io/gorm"
)

func pofi(x int) *int       { return &x }
func pofs(x string) *string { return &x }
func pofb(x bool) *bool     { return &x }

var conn *gorm.DB

func getDBConn() *gorm.DB {
	if conn == nil {
		dbConfig, err := db.NewConfigFromEnv()
		if err != nil {
			panic(fmt.Sprintf("error creating database config: %s", err.Error()))
		}

		conn, err = db.InitDB(dbConfig)
		if err != nil {
			panic(fmt.Sprintf("error initializing database: %s", err.Error()))
		}
	}
	return conn
}

func cleanTables() {
	conn := getDBConn()
	sql := ` TRUNCATE tags, documents CASCADE; `
	conn.Exec(sql)
}

func TestRepositoryGet(t *testing.T) {
	queries := []string{
		`
		INSERT INTO documents (id, created_at, updated_at, title, body)
		VALUES
			( '483b1030-febd-45cc-beca-3d026b8e538c',
				'2021-05-01 12:52:00.000 +00:00',
				'2021-05-01 12:52:00.000 +00:00',
				'myTestTitle',
				'myBeachBody'
			)`,
		`
		INSERT INTO tags ( id, title, created_at, updated_at)
		VALUES
			( '918c6408-a4f5-4bfe-879c-fc1f35cdb82f',
				'tagX',
				'2021-05-01 17:12:18.023142+00',
				'2021-05-01 17:15:18.023142+00'
			),
			( '1d4a172a-55c2-47ab-82af-e03a7050c81d',
				'tagY',
				'2021-05-01 17:12:18.023142+00',
				'2021-05-01 17:15:18.023142+00'
			)
		`,
		`
		INSERT INTO document_tags ( tag_id, document_id)
		VALUES
			( '918c6408-a4f5-4bfe-879c-fc1f35cdb82f',
				'483b1030-febd-45cc-beca-3d026b8e538c'
			),
			( '1d4a172a-55c2-47ab-82af-e03a7050c81d',
				'483b1030-febd-45cc-beca-3d026b8e538c'
			)`,
	}

	conn := getDBConn()
	for _, q := range queries {
		conn.Exec(q)
	}
	defer cleanTables()

	t.Run("test that fetched document is correctly formatted", func(t *testing.T) {
		repo := NewRepository(conn)
		doc, err := repo.Get("483b1030-febd-45cc-beca-3d026b8e538c")
		if err != nil {
			t.Fatalf("got unexpected error: %s", err)
		}

		expectedDoc := db.Document{
			ID:        uuid.MustParse("483b1030-febd-45cc-beca-3d026b8e538c"),
			Title:     "myTestTitle",
			Body:      "myBeachBody",
			CreatedAt: time.Date(2021, time.May, 1, 12, 52, 0, 0, time.UTC),
			UpdatedAt: time.Date(2021, time.May, 1, 12, 52, 0, 0, time.UTC),
			Tags:      make([]db.Tag, 0, 10),
		}
		expectedDoc.Tags = append(expectedDoc.Tags, db.Tag{
			ID:    uuid.MustParse("918c6408-a4f5-4bfe-879c-fc1f35cdb82f"),
			Title: "tagX",
		})
		expectedDoc.Tags = append(expectedDoc.Tags, db.Tag{
			ID:    uuid.MustParse("1d4a172a-55c2-47ab-82af-e03a7050c81d"),
			Title: "tagY",
		})

		if !doc.Equal(&expectedDoc) {
			t.Fatalf("expected doc: '%v', got '%v'", expectedDoc, *doc)
		}
	})
}

func TestRepoNumDocuments(t *testing.T) {
	conn := getDBConn()

	var q string
	for i := 0; i < 73; i++ {
		id := uuid.New()
		q = fmt.Sprintf(`
			INSERT INTO documents (id, created_at, updated_at, title, body)
			VALUES
				( '%s',
					'2021-05-01 12:52:00.000 +00:00',
					'2021-05-01 12:52:00.000 +00:00',
					'myTestTitle',
					'body'
				)`, id)
		conn.Exec(q)
	}
	defer cleanTables()

	repo := NewRepository(conn)

	t.Run("test that correct number of documents is returned", func(t *testing.T) {
		numDocs, err := repo.GetNumDocuments()
		if err != nil {
			t.Fatalf("unable to get num documents: %s", err)
		}

		if numDocs != 73 {
			t.Fatalf("expected 73 documents, got %d", numDocs)
		}
	})
}

func TestRepoGetAll(t *testing.T) {
	conn := getDBConn()

	ids := []interface{}{
		"ca29e280-4abc-488e-bbdd-c2c1d2c57729",
		"6bf519be-0771-41db-8d4e-c79391f0f238",
		"a9a3c4a2-2c39-4729-928a-88b1a201c761",
		"8a663808-acbb-4566-b36b-f995b32cea9d",
	}

	q := fmt.Sprintf(`
		INSERT INTO documents (id, created_at, updated_at, title)
		VALUES
			( '%s', '2021-02-01 12:52:00.000 +00:00', '2021-05-05 12:52:00.000 +00:00', 'title4'),
			( '%s', '2021-03-01 12:52:00.000 +00:00', '2021-05-03 12:52:00.000 +00:00', 'title3'),
			( '%s', '2021-04-01 12:52:00.000 +00:00', '2021-05-04 12:52:00.000 +00:00', 'title2'),
			( '%s', '2021-05-01 12:52:00.000 +00:00', '2021-05-01 12:52:00.000 +00:00', 'title1')
			`, ids...)
	conn.Exec(q)
	defer cleanTables()

	repo := NewRepository(conn)

	tests := []struct {
		name        string
		limit       *int
		offset      *int
		sortBy      *string
		sortDesc    *bool
		expectedIDs []string
	}{
		{
			name:   "test that limit works and default sorting id DESC",
			limit:  pofi(2),
			sortBy: pofs("title"),
			expectedIDs: []string{
				ids[0].(string),
				ids[1].(string),
			},
		},
		{
			name:     "test that sorting works in ASC order",
			limit:    pofi(5),
			sortBy:   pofs("title"),
			sortDesc: pofb(false),
			expectedIDs: []string{
				ids[3].(string),
				ids[2].(string),
				ids[1].(string),
				ids[0].(string),
			},
		},
		{
			name:   "test that offset works with default sorting",
			limit:  pofi(2),
			offset: pofi(1),
			expectedIDs: []string{
				ids[2].(string),
				ids[1].(string),
			},
		},
		{
			name:     "test that offset works with custom sorting",
			limit:    pofi(2),
			sortBy:   pofs("updated_at"),
			offset:   pofi(2),
			sortDesc: pofb(false),
			expectedIDs: []string{
				ids[2].(string),
				ids[0].(string),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := repo.GetAll(test.limit, test.offset, test.sortBy, test.sortDesc)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			resIDs := []string{}
			for i := range res {
				resIDs = append(resIDs, res[i].ID.String())
			}
			if !reflect.DeepEqual(resIDs, test.expectedIDs) {
				t.Fatalf("expected ids: '%v', got '%v'", test.expectedIDs, resIDs)
			}
		})
	}
}

func TestRepoGetNumDocumentsWithTag(t *testing.T) {
	conn := getDBConn()

	docUUIDs := []interface{}{
		"11604a76-ed57-4389-a4bf-f0f20e41b73c",
		"1db47be4-3ab8-446d-a679-e0c6e8873b6c",
		"c86f2aee-240c-4440-b212-52506717cefe",
		"6ecbd1c4-64b0-4a33-9f79-6ac50f3c85bc",
	}
	tagUUIDs := []interface{}{
		"1533c1a6-8c9b-430a-9937-de7ac6e9af57",
		"d2c9e5b4-bbdd-4137-a83c-f487de109207",
		"5cd3c8e6-6a02-40ad-afb0-634e1f41f0d2",
		"bdf64f7f-2433-45be-bf9c-d912dd93d99a",
	}

	queries := []string{
		fmt.Sprintf(`
		INSERT INTO documents (id, created_at, updated_at, title)
		VALUES
			('%s', '2021-05-01 12:52:00.000 +00:00', '2021-05-01 12:52:00.000 +00:00', 'title1'),
			('%s', '2021-05-01 12:52:00.000 +00:00', '2021-05-01 12:52:00.000 +00:00', 'title2'),
			('%s', '2021-05-01 12:52:00.000 +00:00', '2021-05-01 12:52:00.000 +00:00', 'title3'),
			('%s', '2021-05-01 12:52:00.000 +00:00', '2021-05-01 12:52:00.000 +00:00', 'title4')
			`, docUUIDs...),
		fmt.Sprintf(`
		INSERT INTO tags (id, title)
		VALUES
			( '%s', 'tag1'),
			( '%s', 'tag2'),
			( '%s', 'tag3'),
			( '%s', 'tag4')
		`, tagUUIDs...),
		fmt.Sprintf(`
		INSERT INTO document_tags ( tag_id, document_id)
		VALUES
			('%s', '%s'),
			('%s', '%s'),
			('%s', '%s'),
			('%s', '%s'),
			('%s', '%s'),
			('%s', '%s'),
			('%s', '%s')`,
			tagUUIDs[0], docUUIDs[0],
			tagUUIDs[0], docUUIDs[1],
			tagUUIDs[0], docUUIDs[2],
			tagUUIDs[1], docUUIDs[2],
			tagUUIDs[1], docUUIDs[3],
			tagUUIDs[2], docUUIDs[0],
			tagUUIDs[2], docUUIDs[3],
		),
	}
	for _, q := range queries {
		conn.Exec(q)
	}
	defer cleanTables()

	repo := NewRepository(conn)

	tests := []struct {
		name        string
		tagIDs      []string
		expectedNum int64
	}{
		{
			name:        "test that 0 is returned when the tag does not exist",
			tagIDs:      []string{uuid.New().String()},
			expectedNum: 0,
		},
		{
			name:        "test that 0 is returned when tag does not have documents",
			tagIDs:      []string{fmt.Sprintf("%s", tagUUIDs[3])},
			expectedNum: 0,
		},
		{
			name:        "test that tag1 has 3 documents",
			tagIDs:      []string{fmt.Sprintf("%s", tagUUIDs[0])},
			expectedNum: 3,
		},
		{
			name:        "test that tag2 has 2 documents",
			tagIDs:      []string{fmt.Sprintf("%s", tagUUIDs[1])},
			expectedNum: 2,
		},
		{
			name: "test that tag2 and tag3 have 3 unique documents",
			tagIDs: []string{
				fmt.Sprintf("%s", tagUUIDs[1]),
				fmt.Sprintf("%s", tagUUIDs[2]),
			},
			expectedNum: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := repo.GetNumDocumentsWithTag(test.tagIDs)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			if res != test.expectedNum {
				t.Fatalf("expected %d results, but got %d", test.expectedNum, res)
			}
		})
	}
}

func TestRepoGetAllByTag(t *testing.T) {
	conn := getDBConn()

	docUUIDs := []interface{}{
		"11604a76-ed57-4389-a4bf-f0f20e41b73c",
		"1db47be4-3ab8-446d-a679-e0c6e8873b6c",
		"c86f2aee-240c-4440-b212-52506717cefe",
		"6ecbd1c4-64b0-4a33-9f79-6ac50f3c85bc",
	}
	tagUUIDs := []interface{}{
		"1533c1a6-8c9b-430a-9937-de7ac6e9af57",
		"d2c9e5b4-bbdd-4137-a83c-f487de109207",
		"5cd3c8e6-6a02-40ad-afb0-634e1f41f0d2",
		"bdf64f7f-2433-45be-bf9c-d912dd93d99a",
	}

	queries := []string{
		fmt.Sprintf(`
		INSERT INTO documents (id, created_at, updated_at, title)
		VALUES
			( '%s', '2021-02-01 12:52:00.000 +00:00', '2021-05-05 12:52:00.000 +00:00', 'title4'),
			( '%s', '2021-03-01 12:52:00.000 +00:00', '2021-05-03 12:52:00.000 +00:00', 'title3'),
			( '%s', '2021-04-01 12:52:00.000 +00:00', '2021-05-04 12:52:00.000 +00:00', 'title2'),
			( '%s', '2021-05-01 12:52:00.000 +00:00', '2021-05-01 12:52:00.000 +00:00', 'title1')
			`, docUUIDs...),
		fmt.Sprintf(`
		INSERT INTO tags (id, title)
		VALUES
			( '%s', 'tag1'),
			( '%s', 'tag2'),
			( '%s', 'tag3'),
			( '%s', 'tag4')
		`, tagUUIDs...),
		fmt.Sprintf(`
		INSERT INTO document_tags ( tag_id, document_id)
		VALUES
			('%s', '%s'),
			('%s', '%s'),
			('%s', '%s'),
			('%s', '%s'),
			('%s', '%s'),
			('%s', '%s'),
			('%s', '%s')`,
			tagUUIDs[0], docUUIDs[0],
			tagUUIDs[0], docUUIDs[1],
			tagUUIDs[0], docUUIDs[2],
			tagUUIDs[1], docUUIDs[2],
			tagUUIDs[1], docUUIDs[3],
			tagUUIDs[2], docUUIDs[0],
			tagUUIDs[2], docUUIDs[3],
		),
	}
	for _, q := range queries {
		conn.Exec(q)
	}
	defer cleanTables()

	repo := NewRepository(conn)

	tests := []struct {
		name        string
		limit       *int
		offset      *int
		sortBy      *string
		sortDesc    *bool
		tagIDs      []string
		expectedIDs []string
	}{
		{
			name:        "test that no results are returned when no tags are specified",
			expectedIDs: []string{},
		},
		{
			name:        "test that no documents are returned when the tag does not exist",
			tagIDs:      []string{uuid.New().String()},
			expectedIDs: []string{},
		},
		{
			name:        "test that no documents are returned when tag does not have documents",
			tagIDs:      []string{fmt.Sprintf("%s", tagUUIDs[3])},
			expectedIDs: []string{},
		},
		{
			name:   "test that limit is respected when calling tag1 who has 3 documents",
			tagIDs: []string{fmt.Sprintf("%s", tagUUIDs[0])},
			limit:  pofi(2),
			expectedIDs: []string{
				fmt.Sprintf("%s", docUUIDs[2]),
				fmt.Sprintf("%s", docUUIDs[1]),
			},
		},
		{
			name:     "test that sorting works in ASC order",
			limit:    pofi(5),
			sortBy:   pofs("title"),
			sortDesc: pofb(false),
			tagIDs:   []string{fmt.Sprintf("%s", tagUUIDs[0])},
			expectedIDs: []string{
				docUUIDs[2].(string),
				docUUIDs[1].(string),
				docUUIDs[0].(string),
			},
		},
		{
			name:   "test that offset works with default sorting",
			limit:  pofi(2),
			offset: pofi(1),
			tagIDs: []string{fmt.Sprintf("%s", tagUUIDs[0])},
			expectedIDs: []string{
				docUUIDs[1].(string),
				docUUIDs[0].(string),
			},
		},
		{
			name:     "test that offset works with custom sorting",
			limit:    pofi(2),
			sortBy:   pofs("updated_at"),
			sortDesc: pofb(false),
			tagIDs:   []string{fmt.Sprintf("%s", tagUUIDs[1])},
			expectedIDs: []string{
				docUUIDs[3].(string),
				docUUIDs[2].(string),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := repo.GetAllByTag(test.tagIDs, test.limit, test.offset, test.sortBy, test.sortDesc)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			resIDs := []string{}
			for i := range res {
				resIDs = append(resIDs, res[i].ID.String())
			}
			if !reflect.DeepEqual(resIDs, test.expectedIDs) {
				t.Fatalf("expected ids: '%v', got '%v'", test.expectedIDs, resIDs)
			}
		})
	}
}
