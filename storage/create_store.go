package storage

import (
	"github.com/offerni/graphqllearning"
	"github.com/offerni/graphqllearning/mocks"
)

func CreateStore(opts CreateStoreOpts) (*graphqllearning.Store, error) {
	id := graphqllearning.RandomString(10)

	mocks.Stores[id] = &graphqllearning.Store{
		ID:   id,
		Name: opts.Name,
	}

	return mocks.Stores[id], nil
}

type CreateStoreOpts struct {
	Name string
}
