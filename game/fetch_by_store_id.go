package game

import (
	"context"
	"fmt"

	"github.com/offerni/graphqllearning/http"
)

type FetchGameByStoreIDResponse struct {
	Data []*Game
}

func FetchGamesByStoreID(c context.Context, storeID string) (*FetchGameByStoreIDResponse, error) {
	if storeID != http.StorePublicID {
		return nil, fmt.Errorf("NOT FOUND")
	}

	return &FetchGameByStoreIDResponse{
		Data: []*Game{
			{
				ID:      GamePublicID,
				StoreID: http.StorePublicID,
				Name:    "Satisfactory",
				Price:   "19.99",
			},
			{
				ID:      GamePublicID,
				StoreID: http.StorePublicID,
				Name:    "The Wicher",
				Price:   "19.99",
			},
		},
	}, nil
}

type Game struct {
	ID      string
	StoreID string
	Name    string
	Price   string
}
