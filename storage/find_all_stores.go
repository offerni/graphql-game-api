package storage

import (
	"github.com/offerni/graphqllearning"
	"github.com/offerni/graphqllearning/mocks"
)

func FindAllStores() ([]*graphqllearning.Store, error) {
	stores := []*graphqllearning.Store{}
	for _, store := range mocks.Stores {
		stores = append(stores, &graphqllearning.Store{
			ID:   store.ID,
			Name: store.Name,
		})

	}

	return stores, nil
}
