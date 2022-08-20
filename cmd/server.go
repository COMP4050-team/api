package main

import (
	"fmt"
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/COMP4050/square-team-5/api/graph"
	"github.com/COMP4050/square-team-5/api/graph/generated"
	"github.com/COMP4050/square-team-5/api/internal/pkg/config"
	"github.com/COMP4050/square-team-5/api/internal/pkg/db"
	"github.com/COMP4050/square-team-5/api/internal/pkg/web/auth"
)

func main() {
	config := config.NewConfig()

	db := db.NewDB(config.DBFilePath)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db, Config: config}}))

	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://*.vercel.app"},
		AllowCredentials: true,
		AllowHeaders:     []string{"Content-Type"},
	}))
	r.Use(auth.AuthHandler())

	r.Any("/", gin.WrapH(playground.Handler("GraphQL playground", "/query")))
	r.POST("/query", gin.WrapH(srv))

	log.Printf("connect to http://localhost:%d/ for GraphQL playground", config.Port)
	r.Run(fmt.Sprintf(":%d", config.Port))
}
