package game

import (
	"context"
	"fmt"

	"github.com/offerni/graphqllearning"
)

const GamePublicID string = "qwerty"

func Fetch(c context.Context, id string) (*graphqllearning.Game, error) {
	if id != GamePublicID {
		return nil, fmt.Errorf("NOT FOUND")
	}

	return &graphqllearning.Game{
		ID:      GamePublicID,
		StoreID: graphqllearning.StorePublicID,
		Name:    "The Wicher",
		Price:   "19.99",
	}, nil
}
