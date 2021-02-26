package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/akrylysov/algnhsa"
	"github.com/gotha/niuniu-cms/data"
	"github.com/gotha/niuniu-cms/graph"
	"github.com/gotha/niuniu-cms/graph/generated"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const defaultPort = "8080"

func initDB() (*gorm.DB, error) {

	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}
	sslmode := os.Getenv("DB_SSL_MODE")
	if sslmode == "" {
		sslmode = "disable"
	}
	timezone := os.Getenv("DB_TIMEZONE")
	if timezone == "" {
		timezone = "Europe/Sofia"
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host,
		username,
		password,
		name,
		port,
		sslmode,
		timezone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	}

	migrateDB := os.Getenv("MIGRATE_DB")
	if migrateDB == "true" {

		q := `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`
		res := db.Exec(q)
		if res.Error != nil {
			return nil, fmt.Errorf("could not create uuid-ossp extension")
		}

		err = db.AutoMigrate(&data.Document{}, &data.Tag{}, &data.Attachment{})
		if err != nil {
			return nil, fmt.Errorf("could not migrate database schema: %w", err)
		}
	}

	return db, nil
}

func main() {

	db, err := initDB()
	if err != nil {
		fmt.Printf("error initializing database: %s", err.Error())
		os.Exit(1)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	documentService := data.NewDocumentService(db)
	tagService := data.NewTagService(db)
	resolver := graph.NewResolver(
		tagService,
		documentService,
	)
	config := generated.Config{
		Resolvers: resolver,
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	if os.Getenv("LAMBDA") == "true" {
		algnhsa.ListenAndServe(http.DefaultServeMux, nil)
	} else {
		log.Printf("GraphQL playground started at http://localhost:%s/", port)
		log.Fatal(http.ListenAndServe(":"+port, nil))
	}
}
