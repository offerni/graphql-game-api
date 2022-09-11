package http

import (
	"net/http"

	"github.com/labstack/echo"
)

const StorePublicID string = "asdfg"

var stores []map[string]*FetchStoreResponse

func FetchStore(c echo.Context) error {
	id := c.Param("id")
	if id != StorePublicID {
		return c.JSON(http.StatusNotFound, "NO STORE")
	}

	return c.JSON(http.StatusOK, &FetchStoreResponse{
		ID:   StorePublicID,
		Name: "Steam",
	})
}

type FetchStoreResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
