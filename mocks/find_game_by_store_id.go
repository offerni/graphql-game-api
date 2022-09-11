package mocks

import (
	"fmt"

	"github.com/offerni/graphqllearning"
)

// not ideal but just for testing ;)
var gamesByStore map[string][]*graphqllearning.Game

func FindGameByStoreByID(storeID string) ([]*graphqllearning.Game, error) {

	gamesByStore = make(map[string][]*graphqllearning.Game)

	gamesByStore["steam"] = []*graphqllearning.Game{
		{
			ID:      "1",
			StoreID: "steam",
			Name:    "Steam",
			Price:   "19.99",
		},

		{
			ID:      "2",
			StoreID: "steam",
			Name:    "GTA V",
			Price:   "50.21",
		},

		{
			ID:      "3",
			StoreID: "steam",
			Name:    "Subnautica",
			Price:   "10.00",
		},

		{
			ID:      "4",
			StoreID: "epic",
			Name:    "Tony Hawk's Pro Skater",
			Price:   "35.35",
		},
	}

	gamesByStore["epic"] = []*graphqllearning.Game{
		{
			ID:      "4",
			StoreID: "epic",
			Name:    "Tony Hawk's Pro Skater",
			Price:   "35.35",
		},
	}

	if gamesByStore[storeID] == nil {
		return nil, fmt.Errorf("STORE DOESN'T EXIST")
	}

	return gamesByStore[storeID], nil
}
