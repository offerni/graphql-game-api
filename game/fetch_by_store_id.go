package game

import (
	"context"
	"fmt"

	"github.com/offerni/graphqllearning"
	"github.com/offerni/graphqllearning/storage"
)

func FetchGamesByStoreID(c context.Context, storeID string) ([]*graphqllearning.Game, error) {
	if storeID == "" {
		return nil, fmt.Errorf("NOT FOUND")
	}

	games, err := storage.FindGameByStoreByID(storeID)
	if err != nil {
		return nil, fmt.Errorf("NO STORE")
	}

	return games, nil
}
