package store

import (
	"context"
	"fmt"

	"github.com/offerni/graphqllearning"
	"github.com/offerni/graphqllearning/mocks"
)

func Create(c context.Context, opts CreateOpts) (*graphqllearning.Store, error) {
	store, err := mocks.CreateStore(mocks.CreateOpts{
		Name: opts.Name,
	})
	if err != nil {
		return nil, fmt.Errorf("NO STORE")
	}

	return store, nil
}

type CreateOpts struct {
	Name string
}
