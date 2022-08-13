package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/COMP4050/square-team-5/api/graph"
	"github.com/COMP4050/square-team-5/api/graph/generated"
	"github.com/COMP4050/square-team-5/api/internal/pkg/db"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dbFilePath := os.Getenv("DB_FILE_PATH")

	db := db.NewDB(dbFilePath)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	corsHandler := cors.Default().Handler(srv)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", corsHandler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
