package storage

import (
	"fmt"

	"github.com/offerni/graphqllearning"
	"github.com/offerni/graphqllearning/mocks"
)

func FindGameByID(id string) (*graphqllearning.Game, error) {
	if mocks.Games[id] == nil {
		return nil, fmt.Errorf("GAME DOESN'T EXIST")
	}

	return mocks.Games[id], nil
}
