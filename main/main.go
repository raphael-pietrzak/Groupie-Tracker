package main

import (
	"fmt"
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

	fmt.Println("Listening at http://localhost:8000")

	http.ListenAndServe("localhost:8000", nil)
}
