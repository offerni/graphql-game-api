package mocks

import (
	"fmt"

	"github.com/offerni/graphqllearning"
)

func CreateStore(opts CreateOpts) (*graphqllearning.Store, error) {
	id := graphqllearning.RandomString(10)

	if stores[id].ID == id {
		return nil, fmt.Errorf("Store already exists")
	}

	stores[id] = &graphqllearning.Store{
		Name: opts.Name,
	}

	return stores[id], nil
}

type CreateOpts struct {
	Name string
}
