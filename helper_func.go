package main
import(
  "math/rand"
  "github.com/google/uuid"

)

//randomization
func random(min, max int, seedgarb int64) int {
	rand.Seed(seedgarb)
	return rand.Intn(max-min) + min
}

//check if []int contains an int
func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

//returns a UUID string
func UUID() string {
	u, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	return u.String()
}
