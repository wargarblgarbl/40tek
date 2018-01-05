package main

import (
//	"fmt"
	"math/rand"
	"time"
)

func random(min, max int, seedgarb int64) int {
	rand.Seed(seedgarb)
	return rand.Intn(max-min) + min
}

func contains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func PlayPlanets() []Card{
	z := LoadPlanet()
  var sel_planets []Card
  var atlas []int
  aynrand := 99999999
  for uu := 0; uu <= 6; uu++ {
		//fmt.Println(uu)
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
//	fmt.Println(sel_planets)
  return(sel_planets)
}
