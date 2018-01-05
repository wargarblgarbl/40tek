package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

func LoadPlanet() []Card {
	dir := "./cards/Core/"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	var c []Card
	for _, f := range files {
		var planet Card
		raw, err := ioutil.ReadFile(dir + f.Name())
		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal(raw, &planet)
		// Potentially make function generic
		for _, i := range planet.Type {
			if i == "Planet" {
				c = append(c, planet)
			}
		}
	}
	return c
}

func PlayPlanets() []Card {
	z := LoadPlanet()
	var sel_planets []Card
	var atlas []int
	aynrand := 99999999
	for uu := 0; uu <= 6; uu++ {
	Random:
		for {
			aynrand = random(0, len(z)-1, time.Now().UnixNano()/int64(len(z)))
			if !contains(atlas, aynrand) {
				sel_planets = append(sel_planets, z[aynrand])
				atlas = append(atlas, aynrand)
				break Random
			} else {
				continue Random
			}
		}
	}
	return (sel_planets)
}
