package groupie

import (
	"html/template"
	"net/http"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	
	APIRequests()

	tmpl := template.Must(template.ParseFiles("static/index.html"))
	new := ArtistStruct{Tab: ArtistTab,}
	tmpl.Execute(w, new)
}
