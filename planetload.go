package main

import (
	"encoding/json"
	//  "fmt"
	"io/ioutil"
	//  "os"
	"log"
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
