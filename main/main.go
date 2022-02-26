package main

import (
	"fmt"
	g "groupie"
	"net/http"
	"sort"
	"strings"
	// "github.com/gorilla/mux"
)

type Artistes struct {
	a int
	b string
	c []string
}

func main() {

	http.HandleFunc("/", g.MainPage)
	http.HandleFunc("/artist", g.Artiste)
	http.HandleFunc("/filter", g.Filter)

	g.APIRequestsLoc()
	for _, v := range g.LocationTab {
		for _, a := range v.Locations {
			split := strings.Split(a, "-", )
			c := strings.Title(strings.Replace(split[1],"_", " ", -1))
			if g.ContainsCountry(g.Countries,c){
				g.Countries = append(g.Countries, c)
			}
		}
	}
	sort.Strings(g.Countries)

	//Show #CSS
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Listening at http://localhost:5500")
	http.ListenAndServe("localhost:5500", nil)
}
