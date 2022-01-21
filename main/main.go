package main

import "net/http"



func main() {
	http.HandleFunc("/", g.MainPage)
	http.ListenAndServe()
}