package main

import ()

type Playfield struct {
	Planets []Card
	Player  []Player
}

type Player struct {
	ID          string
	Warlord 		[]Card
	Hand        []Card
	PlanetSlots [][]Card
	HQ          []Card
	Graveyard   []Card
}
