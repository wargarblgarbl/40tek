package main

import ()


func CreatePlayer() *Player {
	return &Player{
		ID: UUID(),
	}
}
