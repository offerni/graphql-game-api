package game

import (
	"context"
	"fmt"

	"github.com/offerni/graphqllearning"
	"github.com/offerni/graphqllearning/mocks"
)

func FetchGamesByStoreID(c context.Context, storeID string) ([]*graphqllearning.Game, error) {
	if storeID == "" {
		return nil, fmt.Errorf("NOT FOUND")
	}

	games, err := mocks.FindGameByStoreByID(storeID)
	if err != nil {
		return nil, fmt.Errorf("NO STORE")
	}

	return games, nil
}
