package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/offerni/graphql-learning/graph"
	"github.com/offerni/graphql-learning/graph/generated"
)

const graphQLDefaultPort = "8081"
const apiDefaultPort = "8080"

func main() {
	initGraphQLServer()
}

func initGraphQLServer() {
	port := os.Getenv("GRAPHQL_PORT")
	if port == "" {
		port = graphQLDefaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// func initEchoServer(ctx context.Context) {
// 	port := os.Getenv("API_PORT")
// 	if port == "" {
// 		port = apiDefaultPort
// 	}

// 	e := echo.New()
// 	e.GET("health", apiHttp.Health)

// 	e.GET("games/:id", apiHttp.FetchGame(ctx, "sdfsdf"))

// 	// ...
// 	if err := e.Start(":8080"); err != http.ErrServerClosed {
// 		log.Fatal(err)
// 	}

// 	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
// }
