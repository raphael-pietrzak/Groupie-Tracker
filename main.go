package main

import (
	"fmt"
	g "groupie/groupie"
	"net/http"
)

func main() {

	g.APIRequests()

	http.HandleFunc("/", g.MainPage)
	http.HandleFunc("/artist", g.Artiste)
	http.HandleFunc("/filter", g.Filter)

	//Show #CSS
	fs := http.FileServer(http.Dir("./groupie/static"))
	http.Handle("/groupie/static/", http.StripPrefix("/groupie/static/", fs))

	fmt.Println("Listening at http://localhost:5500")
	http.ListenAndServe("localhost:5500", nil)

}
