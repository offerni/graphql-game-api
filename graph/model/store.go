package model

type Store struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Games []*Game `json:"games"`
}
