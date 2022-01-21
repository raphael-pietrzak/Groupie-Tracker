package main

import "net/http"



func main() {
	http.HandleFunc("/", g.Groupie)
	http.ListenAndServe()
}