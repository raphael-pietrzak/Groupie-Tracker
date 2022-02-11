 package main

import (
	"fmt"
	g "groupie"
	"net/http"
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

	//Show #CSS
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Listening at http://localhost:5500")
	http.ListenAndServe("localhost:5500", nil)
}


// func main() {
// 	r := mux.NewRouter()

// 	r.HandleFunc("/", g.MainPage).Methods("GET")
// 	r.HandleFunc("/artist", g.Artiste).Methods("GET")

// 	//Show #CSS
// 	fs := http.FileServer(http.Dir("./static"))
// 	http.Handle("/static/", http.StripPrefix("/static/", fs))

// 	fmt.Println("Listening at http://localhost:5500")
// 	http.ListenAndServe("localhost:5500", r)
// }
