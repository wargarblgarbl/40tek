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
	html = append(html, "<html><center>")
	for _, element := range state.Planets {
		//html = append(html, "<br />"+"a")
		html = append(html, "<img onmouseover=\"bigImg(this)\" onmouseout=\"normalImg(this)\" src=\""+element.Image+"\" height=\"144\" width=\"200\">")
	}
	html = append(html, "")
	exit := strings.Join(html, `<script>
function bigImg(x) {
    x.style.height = "318px";
    x.style.width = "400px";
}

function normalImg(x) {
    x.style.height = "144px";
    x.style.width = "200px";
}
</script>

`)
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
