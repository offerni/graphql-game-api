package mocks

import (
	"fmt"

	"github.com/offerni/graphqllearning"
)

// not ideal but just for testing ;)
var (
	stores map[string]*graphqllearning.Store
)

func FindStoreByID(id string) (*graphqllearning.Store, error) {

	stores = make(map[string]*graphqllearning.Store)

	stores["steam"] = &graphqllearning.Store{
		ID:   "steam",
		Name: "Steam",
	}

	stores["epic"] = &graphqllearning.Store{
		ID:   "epic",
		Name: "Epic Games",
	}

	stores["psn"] = &graphqllearning.Store{
		ID:   "psn",
		Name: "Playstation Store",
	}

	if stores[id] == nil {
		return nil, fmt.Errorf("STORE DOESN'T EXIST")
	}

	return stores[id], nil
}
