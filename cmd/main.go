package main

import (
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo"
	"github.com/offerni/graphqllearning/graph"
	"github.com/offerni/graphqllearning/graph/generated"
	apiHttp "github.com/offerni/graphqllearning/http"
)

const graphQLDefaultPort = "8081"
const apiDefaultPort = "8080"

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		initRESTServer()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		initGraphQLServer()
	}()

	wg.Wait()
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

func initRESTServer() {
	port := os.Getenv("API_PORT")
	if port == "" {
		port = apiDefaultPort
	}

	e := echo.New()
	e.GET("/", apiHttp.Health)
	e.GET("games/:id", apiHttp.FetchGame)
	e.GET("stores/:id", apiHttp.FetchStore)

	log.Printf("connect to http://localhost:%s/ for the REST API Server", port)
	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}
