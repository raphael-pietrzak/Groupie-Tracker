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

	formattedConcertLocations := make(map[string][]string)
	for k, v := range variable {
		split := strings.Split(k, "-")
		city := strings.Title((strings.Replace(split[0], "_", " ", -1)))
		country := strings.Title((strings.Replace(split[1], "_", " ", -1)))
		formattedConcertLocations[city+", "+country] = v
	}
	new := ArtistStruct{Tab2: formattedConcertLocations}
	tmpl.Execute(w, new)
}

func Search(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	search := r.Form.Get("searchBar")
	fmt.Println(search)
}
