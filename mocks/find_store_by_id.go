package mocks

import (
	"fmt"

	"github.com/offerni/graphqllearning"
)

// not ideal but just for testing ;)
var stores = map[string]*graphqllearning.Store{
	"steam": {
		ID:   "steam",
		Name: "Steam",
	},
	"epic": {
		ID:   "epic",
		Name: "Epic Games",
	},

	"psn": {
		ID:   "psn",
		Name: "Playstation Store",
	},
}

func FindStoreByID(id string) (*graphqllearning.Store, error) {
	if stores[id] == nil {
		return nil, fmt.Errorf("STORE DOESN'T EXIST")
	}

	return stores[id], nil
}
