package store

import (
	"context"
	"fmt"

	"github.com/offerni/graphqllearning"
	"github.com/offerni/graphqllearning/storage"
)

func FetchAll(c context.Context) ([]*graphqllearning.Store, error) {
	stores, err := storage.FindAllStores()
	if err != nil {
		return nil, fmt.Errorf("NO STORE")
	}

	return stores, nil
}
