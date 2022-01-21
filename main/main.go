package main

import (
	g "groupie"
	"net/http"
)

func main() {
	http.HandleFunc("/", g.MainPage)

	http.ListenAndServe("localhost:8000", nil)
}
