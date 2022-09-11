package http

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/offerni/graphqllearning/game"
)

func CreateGame(c echo.Context) (err error) {
	req := new(CreateGameRequest)
	if err = c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "BAD REQUEST")
	}

	game, err := game.Create(c.Request().Context(), game.CreateOpts{
		Name:    req.Name,
		StoreID: req.StoreID,
		Price:   req.Price,
	})

	if err != nil {
		return c.JSON(http.StatusNotFound, "NOT FOUND")
	}

	return c.JSON(http.StatusOK, &CreateGameResponse{
		Data: &FetchGameResponse{
			ID:      game.ID,
			Name:    game.Name,
			StoreID: game.StoreID,
			Price:   game.Price,
		},
	})
}

type CreateGameRequest struct {
	Name    string `json:"name"`
	StoreID string `json:"store_id"`
	Price   string `json:"price"`
}

type CreateGameResponse struct {
	Data *FetchGameResponse
}
