package main

import (
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/labstack/echo"
	"github.com/offerni/graphqllearning/graph"
	"github.com/offerni/graphqllearning/graph/generated"
	apiHttp "github.com/offerni/graphqllearning/http"
	"github.com/rs/cors"
)

const graphQLDefaultPort = "8081"
const httpDefaultPort = "8080"

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
	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)
	port := os.Getenv("GRAPHQL_PORT")
	if port == "" {
		port = graphQLDefaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func initRESTServer() {
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = httpDefaultPort
	}

	e := echo.New()
	setHttpRoutes(e)

	log.Printf("connect to http://localhost:%s/ for the REST API Server", port)
	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}

// TODO extract into a different file
func setHttpRoutes(e *echo.Echo) {
	e.GET("/", apiHttp.Health)

	// Stores
	stores := e.Group("stores")
	stores.GET("", apiHttp.FetchAllStores)
	stores.GET("/:id", apiHttp.FetchStore)
	stores.POST("", apiHttp.CreateStore)

	// Games
	games := e.Group("games")
	games.GET("/:id", apiHttp.FetchGame)
	games.POST("", apiHttp.CreateGame)
}
