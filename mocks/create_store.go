package mocks

import (
	"github.com/offerni/graphqllearning"
)

func CreateStore(opts CreateOpts) (*graphqllearning.Store, error) {
	id := graphqllearning.RandomString(10)

	stores[id] = &graphqllearning.Store{
		ID:   id,
		Name: opts.Name,
	}

	return stores[id], nil
}

type CreateOpts struct {
	Name string
}
