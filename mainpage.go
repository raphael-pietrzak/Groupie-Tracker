package groupie

import (
	"html/template"
	"net/http"
)


func MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/index.html"))
	new := 
	new := Artistes{a: 3,b: "rjfn", c:[]string{"fzer"}}
	tmpl.Execute(w, new)

}