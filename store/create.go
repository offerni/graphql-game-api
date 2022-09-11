package store

import (
	"context"
	"fmt"

	"github.com/offerni/graphqllearning"
	storage "github.com/offerni/graphqllearning/storage"
)

func Create(c context.Context, opts CreateOpts) (*graphqllearning.Store, error) {
	store, err := storage.CreateStore(storage.CreateStoreOpts{
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
