package store

import (
	"context"
	"fmt"

	"github.com/offerni/graphqllearning"
	"github.com/offerni/graphqllearning/storage"
)

func Fetch(c context.Context, id string) (*graphqllearning.Store, error) {
	if id == "" {
		return nil, fmt.Errorf("NO STORE")
	}

	store, err := storage.FindStoreByID(id)
	if err != nil {
		return nil, fmt.Errorf("NO STORE")
	}

	return store, nil
}
