package store

import (
	"context"
	"fmt"
)

const StorePublicID string = "asdfg"

type Store struct {
	ID   string
	Name string
}

func Fetch(c context.Context, id string) (*Store, error) {
	if id != StorePublicID {
		return nil, fmt.Errorf("NO STORE")
	}

	return &Store{
		ID:   StorePublicID,
		Name: "Steam",
	}, nil
}
