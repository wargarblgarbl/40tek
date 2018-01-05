package main
import (
  "encoding/json"
	"net/http"
  "fmt"
)

func main(){
  http.HandleFunc("/", mainHandler)
//  what := PlayPlanets()
http.ListenAndServe(":8080", nil)


}

func PlayInit(output chan *Playfield) {
  b := &Playfield{
    Planets: PlayPlanets(),
  }
  //Flip the 7 planets
  for uu := 0; uu <= 4; uu++ {
    b.Planets[uu].FlipUp()
  }

output <- b
close(output)

}




func mainHandler(w http.ResponseWriter, r *http.Request){
  var playchan = make(chan *Playfield)
  go PlayInit(playchan)
  b := <- playchan
  //b is our actual playfield. Let's obfuscate it. 
  PlanetObfuscator(b)
 a, err := json.Marshal(b)
 if err != nil {
   fmt.Println(err)
 }
  w.Write(a)
}

