package http

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/offerni/graphqllearning"
	"github.com/offerni/graphqllearning/game"
)

func FetchGamesByStoreID(c echo.Context, storeID string) error {
	if storeID != graphqllearning.StorePublicID {
		return c.JSON(http.StatusUnprocessableEntity, "STORE ID IS REQUIRED")
	}

	resp, err := game.FetchGamesByStoreID(c.Request().Context(), storeID)
	if err != nil {
		return c.JSON(http.StatusNotFound, "NOT FOUND")
	}

	games := []*FetchGameResponse{}
	for _, game := range resp {
		games = append(games, &FetchGameResponse{
			ID:      game.ID,
			StoreID: game.StoreID,
			Name:    game.Name,
			Price:   game.Price,
		})
	}

	return c.JSON(http.StatusOK, games)

}
