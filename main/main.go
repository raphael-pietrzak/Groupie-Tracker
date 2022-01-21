package main

import (
	g "groupie"
	"net/http"
)

func main() {
	http.HandleFunc("/", g.Groupie)

	http.ListenAndServe("localhost:8080", nil)
}
