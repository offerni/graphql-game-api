package http

import (
	"context"
	"fmt"
)

const GamePublicID string = "qwerty"

var games []map[string]*FetchGameResponse

type FetchGameResponse struct {
	ID      string `json:"id"`
	StoreID string `json:"store_id"`
	Name    string `json:"name"`
	Price   string `json:"price"`
}

func FetchGame(c context.Context, id string) (*FetchGameResponse, error) {
	if id != GamePublicID {
		return nil, fmt.Errorf("NOT FOUND")
	}

	return &FetchGameResponse{
		ID:      GamePublicID,
		StoreID: StorePublicID,
		Name:    "The Wicher",
		Price:   "19.99",
	}, nil
}
