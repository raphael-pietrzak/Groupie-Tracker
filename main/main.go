package main

import (
	"fmt"
	g "groupie"
	"net/http"
)

type Artistes struct {
	a int
	b string
	c []string
}

func main() {
	http.HandleFunc("/", g.MainPage)
	http.HandleFunc("/artist", g.Artiste)

	//Show #CSS
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Listening at http://localhost:8000")

	http.ListenAndServe("localhost:8000", nil)
}
