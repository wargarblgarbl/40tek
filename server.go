package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", mainHandler)
	//  what := PlayPlanets()
	http.ListenAndServe(":8080", nil)

}

func PlayInit(output chan *Playfield) {
	var players []Player
	p1 := CreatePlayer()
	p2 := CreatePlayer()
	players = append(players, *p1)
	players = append(players, *p2)

	b := &Playfield{
		Planets: PlayPlanets(),
		Player: players,
	}
	//Flip the 7 planets
	for uu := 0; uu <= 4; uu++ {
		b.Planets[uu].FlipUp()
	}


	output <- b
	close(output)

}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	var playchan = make(chan *Playfield)
	go PlayInit(playchan)
	b := <-playchan
	//b is our actual playfield. Let's obfuscate it.
	PlanetObfuscator(b)
	SetPlanetInitPos(b)
	a, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(a)
}
