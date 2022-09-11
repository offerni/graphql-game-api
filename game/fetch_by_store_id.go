package game

import (
	"context"
	"fmt"

	"github.com/offerni/graphqllearning"
)

func FetchGamesByStoreID(c context.Context, storeID string) ([]*graphqllearning.Game, error) {
	if storeID != graphqllearning.StorePublicID {
		return nil, fmt.Errorf("NOT FOUND")
	}

	return []*graphqllearning.Game{
			{
				ID:      GamePublicID,
				StoreID: graphqllearning.StorePublicID,
				Name:    "Satisfactory",
				Price:   "19.99",
			},
			{
				ID:      GamePublicID,
				StoreID: graphqllearning.StorePublicID,
				Name:    "The Wicher",
				Price:   "19.99",
			},
		},
		nil
}
