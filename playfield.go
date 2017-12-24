package main

import (
	"fmt"
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

func main() {
	z := Load_Planet()
	//fmt.Println(z)
	//myrand := random(0, 9)
	//fmt.Println(z[myrand])
  var sel_planets [7]Card
  var atlas []int
  aynrand := 99999999
  for uu, _ := range sel_planets{
    Random:
    for {
      aynrand = random(0, len(z)-1, time.Now().UnixNano()/int64(len(z)))
      if !contains(atlas, aynrand) {
        atlas = append(atlas, aynrand)
//        fmt.Println(atlas, aynrand, len(z))
        break Random
      } else {
        continue Random
      }
    }
    if uu < 5 {
    z[aynrand].Faceup = true
    } else {
    z[aynrand].Faceup = false
      }
    sel_planets[uu] = z[aynrand]

  }
  //fmt.Println(sel_planets)
	//var play Playfield
	b := &Playfield{	}
	for _, planet  := range sel_planets {
		b.Planets = append(b.Planets, planet)
	}
	fmt.Println(b)
}
