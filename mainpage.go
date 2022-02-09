package groupie

import (
	// "fmt"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("tmpl/index.html"))
	APIRequests()
	new := ArtistStruct{Tab: ArtistTab}
	tmpl.Execute(w, new)
}

func Artiste(w http.ResponseWriter, r *http.Request) {

	// params := mux.Vars(r)
	// fmt.Println(params)
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	link_loc := r.Form.Get("w")

	tmpl := template.Must(template.ParseFiles("tmpl/artist.html"))
	APIRequests2(link_loc) 

	// CiDa := [][]string{}
	// aa := []string{}
	// for key, v := range variable {
	// 	aa = []string{}
	// 	aa = append(aa, key)
	// 	CiDa = append(CiDa, aa)
	// 	CiDa = append(CiDa, v)
	// 	CiDa = append(CiDa, []string{" "})
	// }
	// // var artistConcert map[string][]string
	// for i := 0; i < len(CiDa); i++ {
	// 	for j := len(CiDa[i])-1; j >= 0 ; j-- {
	// 		aa = append(aa, CiDa[i][j])
	// 	}
	// }
	
	formattedConcertLocations := make(map[string][]string) 
	for k, v := range variable{
		split := strings.Split(k,"-")
		city := strings.Title((strings.Replace(split[0],"_"," ",-1))) 
		country := strings.Title((strings.Replace(split[1],"_"," ",-1))) 
		formattedConcertLocations[city + ", " + country] = v 
		fmt.Println(k)

	}
	new := ArtistStruct{Tab2: formattedConcertLocations}
	tmpl.Execute(w, new)
	fmt.Println(variable)
}

// func Location(w http.ResponseWriter, r *http.Request) {
// 	if err := r.ParseForm(); err != nil {
// 		fmt.Fprintf(w, "ParseForm() err: %v", err)
// 		return
// 	}
// 	link_loc := r.Form.Get("w")

// 	tmpl := template.Must(template.ParseFiles("tmpl/artist.html"))
// 	APIRequests2(link_loc)

// 	new := ArtistStruct{Tab3: LocationsTab}
// 	tmpl.Execute(w, new)

// }

// r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	title := vars["title"]
// 	page := vars["page"]
// 	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
// })
// http.ListenAndServe(":80", r)
// }
