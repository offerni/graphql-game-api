package http

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/offerni/graphqllearning/game"
)

var games []map[string]*FetchGameResponse

func FetchGame(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusUnprocessableEntity, "ID IS REQUIRED")
	}

	game, err := game.Fetch(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "NOT FOUND")
	}

	return c.JSON(http.StatusOK, &FetchGameResponse{
		ID:      game.ID,
		StoreID: game.StoreID,
		Name:    game.Name,
		Price:   game.Price,
	})
}

type FetchGameResponse struct {
	ID      string `json:"id"`
	StoreID string `json:"store_id"`
	Name    string `json:"name"`
	Price   string `json:"price"`
}
