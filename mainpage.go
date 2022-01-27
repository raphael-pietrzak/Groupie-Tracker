package groupie

import (
	"html/template"
	"net/http"
)

func MainPage(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("tmpl/index.html"))
	APIRequests()
	new := ArtistStruct{Tab: ArtistTab}
	tmpl.Execute(w, new)

}
