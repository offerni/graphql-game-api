package storage

import (
	"fmt"

	"github.com/offerni/graphqllearning"
	"github.com/offerni/graphqllearning/mocks"
)

func FindGameByStoreByID(storeID string) ([]*graphqllearning.Game, error) {
	if mocks.GamesByStore[storeID] == nil {
		return nil, fmt.Errorf("STORE DOESN'T EXIST")
	}

	return mocks.GamesByStore[storeID], nil
}
