package game

import (
	"context"
	"fmt"

	"github.com/offerni/graphqllearning"
	"github.com/offerni/graphqllearning/storage"
)

func Create(c context.Context, opts CreateOpts) (*graphqllearning.Game, error) {
	game, err := storage.CreateGame(storage.CreateGameOpts{
		Name:    opts.Name,
		StoreID: opts.StoreID,
		Price:   opts.Price,
	})
	if err != nil {
		return nil, fmt.Errorf("NO GAME")
	}

	return game, nil
}

type CreateOpts struct {
	Name    string
	StoreID string
	Price   string
}
