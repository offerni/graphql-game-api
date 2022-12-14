package game

import (
	"context"
	"fmt"

	"github.com/offerni/graphqllearning"
	"github.com/offerni/graphqllearning/storage"
)

func Fetch(c context.Context, id string) (*graphqllearning.Game, error) {
	if id == "" {
		return nil, fmt.Errorf("NOT FOUND")
	}

	game, err := storage.FindGameByID(id)
	if err != nil {
		return nil, fmt.Errorf("NO STORE")
	}

	return game, nil
}
