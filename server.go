package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/Nishi05/graphql-sample/graph"
	"github.com/Nishi05/graphql-sample/graph/services"
	"github.com/Nishi05/graphql-sample/internal"
	_ "github.com/mattn/go-sqlite3"
)

const (
	defaultPort = "8080"
	dbFile      = "./mygraphql.db"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	boil.DebugMode = true

	db, err := sql.Open("sqlite3", fmt.Sprintf("%s?_foreign_keys=on", dbFile))
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	service := services.New(db)

	srv := handler.NewDefaultServer(internal.NewExecutableSchema(internal.Config{Resolvers: &graph.Resolver{
		Srv:     service,
		Loaders: graph.NewLoaders(service),
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
