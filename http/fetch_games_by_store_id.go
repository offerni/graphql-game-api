package http

import (
	"context"
	"fmt"
)

type FetchGameByStoreIDResponse struct {
	Data []*FetchGameResponse
}

func FetchGamesByStoreID(c context.Context, storeID string) (*FetchGameByStoreIDResponse, error) {
	if storeID != StorePublicID {
		return nil, fmt.Errorf("NOT FOUND")
	}

	return &FetchGameByStoreIDResponse{
		Data: []*FetchGameResponse{
			{
				ID:      GamePublicID,
				StoreID: StorePublicID,
				Name:    "Satisfactory",
				Price:   "19.99",
			},
			{
				ID:      GamePublicID,
				StoreID: StorePublicID,
				Name:    "The Wicher",
				Price:   "19.99",
			},
		},
	}, nil
}
