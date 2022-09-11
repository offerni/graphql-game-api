package storage

import (
	"fmt"

	"github.com/offerni/graphqllearning"
	"github.com/offerni/graphqllearning/mocks"
)

func FindStoreByID(id string) (*graphqllearning.Store, error) {
	if mocks.Stores[id] == nil {
		return nil, fmt.Errorf("STORE DOESN'T EXIST")
	}

	return mocks.Stores[id], nil
}
