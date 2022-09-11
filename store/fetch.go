package store

import (
	"context"
	"fmt"

	"github.com/offerni/graphqllearning"
)

const StorePublicID string = "asdfg"

func Fetch(c context.Context, id string) (*graphqllearning.Store, error) {
	if id != StorePublicID {
		return nil, fmt.Errorf("NO STORE")
	}

	return &graphqllearning.Store{
		ID:   StorePublicID,
		Name: "Steam",
	}, nil
}
