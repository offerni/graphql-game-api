package http

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/offerni/graphqllearning/store"
)

var stores []map[string]*FetchStoreResponse

func FetchStore(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusUnprocessableEntity, "ID IS REQUIRED")
	}

	store, err := store.Fetch(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "NOT FOUND")
	}

	return c.JSON(http.StatusOK, &FetchStoreResponse{
		ID:   store.ID,
		Name: store.Name,
	})
}

type FetchStoreResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
