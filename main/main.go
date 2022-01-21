package main

import (
	g "groupie"
	"net/http"
)

type Artistes struct{
	a int
	b string
	c []string
}

func main() {
	http.HandleFunc("/", g.MainPage)

	http.ListenAndServe("localhost:8000", nil)
}
