package http

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/offerni/graphqllearning/store"
)

func CreateStore(c echo.Context) (err error) {
	req := new(CreateStoreRequest)
	if err = c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "BAD REQUEST")
	}

	store, err := store.Create(c.Request().Context(), store.CreateOpts{
		Name: req.Name,
	})

	if err != nil {
		return c.JSON(http.StatusNotFound, "NOT FOUND")
	}

	return c.JSON(http.StatusOK, &CreateStoreResponse{
		Data: &FetchStoreResponse{
			ID:   store.ID,
			Name: store.Name,
		},
	})
}

type CreateStoreRequest struct {
	Name string `json:"name"`
}

type CreateStoreResponse struct {
	Data *FetchStoreResponse
}
