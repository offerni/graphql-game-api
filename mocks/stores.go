package mocks

import "github.com/offerni/graphqllearning"

var Stores = map[string]*graphqllearning.Store{
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
