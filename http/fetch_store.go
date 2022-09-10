package http

import (
	"context"
	"fmt"
)

const StorePublicID string = "asdfg"

var stores []map[string]*FetchStoreResponse

type FetchStoreResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func FetchStore(c context.Context, id string) (*FetchStoreResponse, error) {
	if id != StorePublicID {
		return nil, fmt.Errorf("NO STORE")
	}

	return &FetchStoreResponse{
		ID:   StorePublicID,
		Name: "Steam",
	}, nil
}
