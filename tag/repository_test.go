package tag

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/gotha/niuniu-cms/db"
	"gorm.io/gorm"
)

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
	sql := ` TRUNCATE tags CASCADE; `
	conn.Exec(sql)
}

func TestRepoGetAll(t *testing.T) {
	defer cleanTables()
	conn := getDBConn()
	conn.Exec(`
		INSERT INTO tags (id, title)
		VALUES
			( '918c6408-a4f5-4bfe-879c-fc1f35cdb82f', 'tagX'),
			( '4fef9f0f-42e0-41e8-8917-4effce834aa7', 'tagY'),
			( '1d4a172a-55c2-47ab-82af-e03a7050c81d', 'tagZ')
	`)

	repo := NewRepository(conn)
	res, err := repo.GetAll()
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if len(res) != 3 {
		t.Fatalf("expected 3 tags, but got %d", len(res))
	}
}

func TestRepoGetMultiple(t *testing.T) {
	defer cleanTables()
	conn := getDBConn()
	conn.Exec(`
		INSERT INTO tags (id, title)
		VALUES
			( '918c6408-a4f5-4bfe-879c-fc1f35cdb82f', 'tagX'),
			( '4fef9f0f-42e0-41e8-8917-4effce834aa7', 'tagY'),
			( '1d4a172a-55c2-47ab-82af-e03a7050c81d', 'tagZ')
	`)

	repo := NewRepository(conn)
	res, err := repo.GetMultiple([]string{
		"918c6408-a4f5-4bfe-879c-fc1f35cdb82f",
		"1d4a172a-55c2-47ab-82af-e03a7050c81d",
	})
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if len(res) != 2 {
		t.Fatalf("expected 2 tags, but got %d", len(res))
	}
}

func TestRepoGetTagByTitle(t *testing.T) {
	defer cleanTables()
	conn := getDBConn()
	conn.Exec(`
		INSERT INTO tags (id, title)
		VALUES
			( '918c6408-a4f5-4bfe-879c-fc1f35cdb82f', 'tagX'),
			( '4fef9f0f-42e0-41e8-8917-4effce834aa7', 'tagY'),
			( '1d4a172a-55c2-47ab-82af-e03a7050c81d', 'tagZ')
	`)

	tests := []struct {
		name  string
		title string
		res   *db.Tag
	}{
		{
			name:  "test that tag is returned when it exists",
			title: "tagZ",
			res: &db.Tag{
				ID:    uuid.MustParse("1d4a172a-55c2-47ab-82af-e03a7050c81d"),
				Title: "tagZ",
			},
		},
		{
			name:  "test that nil is returned when tag does not exist",
			title: "tagA",
		},
	}

	repo := NewRepository(conn)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := repo.GetTagByTitle(test.title)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			if res != nil {
				if !res.Equal(test.res) {
					t.Fatalf("expected tag: '%v', got '%v'", test.res, res)
				}

				if test.res == nil {
					t.Fatalf("test did not expect result, but got '%v'", res)
				}
			}
		})
	}
}

func TestRepoGet(t *testing.T) {
	defer cleanTables()
	conn := getDBConn()
	conn.Exec(`
		INSERT INTO tags (id, title)
		VALUES
			( '918c6408-a4f5-4bfe-879c-fc1f35cdb82f', 'tagX'),
			( '4fef9f0f-42e0-41e8-8917-4effce834aa7', 'tagY'),
			( '1d4a172a-55c2-47ab-82af-e03a7050c81d', 'tagZ')
	`)

	tests := []struct {
		name string
		id   string
		res  *db.Tag
	}{
		{
			name: "test that tag is returned when it exists",
			id:   "918c6408-a4f5-4bfe-879c-fc1f35cdb82f",
			res: &db.Tag{
				ID:    uuid.MustParse("918c6408-a4f5-4bfe-879c-fc1f35cdb82f"),
				Title: "tagX",
			},
		},
		{
			name: "test that nil is returned when tag does not exist",
			id:   "aa2777d2-ba85-4217-a490-65d284a5802e",
		},
	}

	repo := NewRepository(conn)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := repo.Get(test.id)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			if res != nil {
				if !res.Equal(test.res) {
					t.Fatalf("expected tag: '%v', got '%v'", test.res, res)
				}

				if test.res == nil {
					t.Fatalf("test did not expect result, but got '%v'", res)
				}
			}
		})
	}
}

func TestRepoCreate(t *testing.T) {
	defer cleanTables()
	conn := getDBConn()
	repo := NewRepository(conn)
	res, err := repo.Create("myTag")
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}

	row := conn.Raw(fmt.Sprintf("SELECT COUNT(*) FROM tags WHERE id = '%s'", res.ID.String())).Row()
	var num int
	row.Scan(&num)
	if num != 1 {
		t.Fatalf("expected tag to be saved, but could not be found in database")
	}
}

func TestRepoUpdate(t *testing.T) {
	defer cleanTables()
	conn := getDBConn()

	repo := NewRepository(conn)
	res, err := repo.Create("myTag")
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}

	_, err = repo.Update(*res, "myUpdatedTag")
	if err != nil {
		t.Fatalf("unexpected err from update: %s", err)
	}

	row := conn.Raw(fmt.Sprintf("SELECT title FROM tags WHERE id = '%s'", res.ID.String())).Row()
	var title string
	row.Scan(&title)
	if title != "myUpdatedTag" {
		t.Fatalf("expected tag title to be update, but could not confirm")
	}
}

func TestRepoDelete(t *testing.T) {
	defer cleanTables()
	conn := getDBConn()
	conn.Exec(`
		INSERT INTO tags (id, title)
		VALUES
			( '918c6408-a4f5-4bfe-879c-fc1f35cdb82f', 'tagX'),
			( '4fef9f0f-42e0-41e8-8917-4effce834aa7', 'tagY'),
			( '1d4a172a-55c2-47ab-82af-e03a7050c81d', 'tagZ')
	`)

	repo := NewRepository(conn)

	err := repo.Delete("4fef9f0f-42e0-41e8-8917-4effce834aa7")
	if err != nil {
		t.Fatalf("unexpected err from update: %s", err)
	}

	row := conn.Raw("SELECT COUNT(*) FROM tags WHERE deleted_at IS NULL").Row()
	var num int
	row.Scan(&num)
	if num != 2 {
		t.Fatalf("expected to have 2 remainig tag, but got %d", num)
	}
}
