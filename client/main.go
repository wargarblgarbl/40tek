package main 
import (	
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
)

func main() {
	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(":8081", nil)

}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	var state Playfield
	jsonstring := curlJson("http://localhost:8080")
	err := json.Unmarshal(jsonstring, &state)
	if err != nil {

	}
	var html []string
	fmt.Println(html)
	html = append(html, "<html>")
	for _, element := range state.Planets {
		//html = append(html, "<br />"+"a")
		if element.Image == "" {
			element.Image = "https://1d4chan.org/images/thumb/4/49/Angry_Marine_Desktop.jpg/800px-Angry_Marine_Desktop.jpg"
		}
		html = append(html, "<img src=\""+element.Image+"\">")
	}
	exit := strings.Join(html, "")
	w.Write([]byte(exit))

}

func curlJson(url string) (stuff []byte){
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	} else {
		defer r.Body.Close()
		contents, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		stuff :=  contents
		return stuff
	}
}
