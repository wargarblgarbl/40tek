package main

type Card struct {
	Tapped       bool
	Faceup       bool
	Position     int
	Damage       int
	Action       Action
	Commadicons  int
	Cost         int
	Draw         int
	Faction      string
	HP           int
	Loyal        bool
	Loyalty      bool
	Name         string
	Power        int
	Buff         int
	Planettype   []string
	Resource     int
	Shield       int
	Signature    bool
	SignatureNum int
	Traits       []string
	Type         []string
}

type Action struct {
	Action []string
	Target []string
}
