package http

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/offerni/graphqllearning/store"
)

func FetchAllStores(c echo.Context) error {
	stores, err := store.FetchAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusNotFound, "NOT FOUND")
	}

	return c.JSON(http.StatusOK, stores)
}

type FetchAllStoresResponse struct {
	Data []*FetchStoreResponse `json:"data"`
}
