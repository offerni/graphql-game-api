package mocks

import (
	"fmt"

	"github.com/offerni/graphqllearning"
)

// not ideal but just for testing ;)
var games = map[string]*graphqllearning.Game{
	"1": {
		ID:      "1",
		StoreID: "steam",
		Name:    "The Witcher",
		Price:   "19.99",
	},

	"2": {
		ID:      "2",
		StoreID: "steam",
		Name:    "GTA V",
		Price:   "50.21",
	},

	"3": {
		ID:      "3",
		StoreID: "steam",
		Name:    "Subnautica",
		Price:   "10.00",
	},

	"4": {
		ID:      "4",
		StoreID: "epic",
		Name:    "Tony Hawk's Pro Skater",
		Price:   "35.35",
	},
}

func FindGameByID(id string) (*graphqllearning.Game, error) {
	if games[id] == nil {
		return nil, fmt.Errorf("GAME DOESN'T EXIST")
	}

	return games[id], nil
}
