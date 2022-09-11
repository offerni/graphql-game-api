package storage

import (
	"github.com/offerni/graphqllearning"
	"github.com/offerni/graphqllearning/mocks"
)

func CreateGame(opts CreateGameOpts) (*graphqllearning.Game, error) {
	id := graphqllearning.RandomString(10)

	mocks.Games[id] = &graphqllearning.Game{
		ID:      id,
		Name:    opts.Name,
		StoreID: opts.StoreID,
		Price:   opts.Price,
	}

	// When creating a game, this relationship collection also needs to be updated
	mocks.GamesByStore[opts.StoreID] = append(mocks.GamesByStore[opts.StoreID], &graphqllearning.Game{
		ID:      id,
		Name:    opts.Name,
		StoreID: opts.StoreID,
		Price:   opts.Price,
	})

	return mocks.Games[id], nil
}

type CreateGameOpts struct {
	Name    string
	StoreID string
	Price   string
}
