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
	"github.com/gotha/niuniu-cms/db"
	"github.com/gotha/niuniu-cms/graph"
	"github.com/gotha/niuniu-cms/graph/generated"
	"github.com/rs/cors"
)

func main() {

	config, err := NewConfigFromEnv()
	if err != nil {
		fmt.Printf("error loading config: %s", err.Error())
		os.Exit(1)
	}

	dbConfig, err := db.NewConfigFromEnv()
	if err != nil {
		fmt.Printf("error creating database config: %s", err.Error())
		os.Exit(1)
	}

	db, err := db.InitDB(dbConfig)
	if err != nil {
		fmt.Printf("error initializing database: %s", err.Error())
		os.Exit(1)
	}

	documentService := data.NewDocumentService(db)
	tagService := data.NewTagService(db)
	attachmentService := data.NewAttachmentService(db)
	resolver := graph.NewResolver(
		tagService,
		documentService,
		attachmentService,
	)
	gqlconfig := generated.Config{
		Resolvers: resolver,
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(gqlconfig))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	http.Handle("/query", cors.Default().Handler(srv))

	if os.Getenv("LAMBDA") == "true" {
		algnhsa.ListenAndServe(http.DefaultServeMux, nil)
	} else {
		log.Printf("GraphQL playground started at http://localhost:%d/", config.Port)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil))
	}
}
