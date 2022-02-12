package groupie

import (
	// "fmt"

	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("tmpl/index.html"))
	APIRequests()
	new := ArtistStruct{Tab: ArtistTab}
	tmpl.Execute(w, new)
}

func Artiste(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	link_loc := r.Form.Get("w")

	tmpl := template.Must(template.ParseFiles("tmpl/artist.html"))

	APIRequests2(link_loc)

	newLocs := make(map[string][]string)
	for k, v := range Relation.DatesLoc {
		fmt.Println(k, v)
		newLocs[k] = v
	}
	for k, v := range newLocs {
		split := strings.Split(k, "-")
		city := strings.Title((strings.Replace(split[0], "_", " ", -1)))
		country := strings.Title((strings.Replace(split[1], "_", " ", -1)))
		delete(Relation.DatesLoc, k)
		Relation.DatesLoc[city+", "+country] = v
	}
	

	tmpl.Execute(w, ArtistStruct{S1: Relation, S2: Artists})
}
