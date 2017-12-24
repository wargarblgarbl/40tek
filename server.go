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
output <- b
close(output)

}

func mainHandler(w http.ResponseWriter, r *http.Request){
  var playchan = make(chan *Playfield)
  go PlayInit(playchan)
  b := <- playchan
 a, err := json.Marshal(b)
 if err != nil {
   fmt.Println(err)
 }
  w.Write(a)
}
