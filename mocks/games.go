package mocks

import "github.com/offerni/graphqllearning"

var Games = map[string]*graphqllearning.Game{
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
