package game

import (
	"context"
	"fmt"

	"github.com/offerni/graphqllearning/http"
)

const GamePublicID string = "qwerty"

func Fetch(c context.Context, id string) (*Game, error) {
	if id != GamePublicID {
		return nil, fmt.Errorf("NOT FOUND")
	}

	return &Game{
		ID:      GamePublicID,
		StoreID: http.StorePublicID,
		Name:    "The Wicher",
		Price:   "19.99",
	}, nil
}
