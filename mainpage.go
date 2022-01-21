package groupie

import (
	"html/template"
	"net/http"
)

type Artistes struct{
	
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/index.html"))
	data := ArtistStruct{
		Tab: ArtistsTab,
	}
	tmpl.Execute(w, data)
}
