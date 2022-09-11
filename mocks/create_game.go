package mocks

import (
	"github.com/offerni/graphqllearning"
)

func CreateGame(opts CreateGameOpts) (*graphqllearning.Game, error) {
	id := graphqllearning.RandomString(10)

	games[id] = &graphqllearning.Game{
		ID:      id,
		Name:    opts.Name,
		StoreID: opts.StoreID,
		Price:   opts.Price,
	}

	return games[id], nil
}

type CreateGameOpts struct {
	Name    string
	StoreID string
	Price   string
}
