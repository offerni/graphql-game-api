package http

import (
	"net/http"

	"github.com/labstack/echo"
)

const GamePublicID string = "qwerty"

var games []map[string]*FetchGameResponse

func FetchGame(c echo.Context) error {
	id := c.Param("id")
	if id != GamePublicID {
		return c.JSON(http.StatusNotFound, "NOT FOUND")
	}

	return c.JSON(http.StatusOK, &FetchGameResponse{
		ID:      GamePublicID,
		StoreID: StorePublicID,
		Name:    "The Wicher",
		Price:   "19.99",
	})
}

type FetchGameResponse struct {
	ID      string `json:"id"`
	StoreID string `json:"store_id"`
	Name    string `json:"name"`
	Price   string `json:"price"`
}
