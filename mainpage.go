package groupie

import (
	"html/template"
	"net/http"
)

type Artistes struct{
	a int
	b string
	c []string
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/index.html"))
	new := Artistes{a: 3,b: "rjfn", c: ["kjsdfh"]}
	tmpl.Execute(w, new)
}
