package main

import "net/http"



func main() {
	http.HandleFunc("/", Groupie)
	http.ListenAndServe("http")
}