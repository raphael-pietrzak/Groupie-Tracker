package groupie

import (
	"html/template"
	"net/http"
)


func MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/index.html"))
<<<<<<< HEAD
	new = Artistes{a: 3,b: "rjfn", c:[]string{"fzer"}}
	tmpl.Execute(w, new)

}
=======
	data := ArtistStruct{
		Tab: ArtistsTab,
	}
	tmpl.Execute(w, data)
}
>>>>>>> dde08957873d2264871dfd8550b7864d944c7c22
