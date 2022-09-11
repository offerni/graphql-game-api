package mocks

import (
	"fmt"

	"github.com/offerni/graphqllearning"
)

// not ideal but just for testing ;)
var games map[string]*graphqllearning.Game

func FindGameByID(id string) (*graphqllearning.Game, error) {
	games = make(map[string]*graphqllearning.Game)

	games["1"] = &graphqllearning.Game{
		ID:      "1",
		StoreID: "steam",
		Name:    "The Witcher",
		Price:   "19.99",
	}

	games["2"] = &graphqllearning.Game{
		ID:      "2",
		StoreID: "steam",
		Name:    "GTA V",
		Price:   "50.21",
	}

	games["3"] = &graphqllearning.Game{
		ID:      "3",
		StoreID: "steam",
		Name:    "Subnautica",
		Price:   "10.00",
	}

	games["4"] = &graphqllearning.Game{
		ID:      "4",
		StoreID: "epic",
		Name:    "Tony Hawk's Pro Skater",
		Price:   "35.35",
	}

	if games[id] == nil {
		return nil, fmt.Errorf("STORE DOESN'T EXIST")
	}

	return games[id], nil
}
