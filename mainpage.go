package groupie

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func MainPage(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("tmpl/index.html"))

	new := ArtistStruct{Tab: ArtistTab, Country: CountryList}
	tmpl.Execute(w, new)
}

func Artiste(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("tmpl/artist.html"))

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	Id_Artit, _ := strconv.Atoi(r.Form.Get("w"))

	tmpl.Execute(w, ArtistTab[Id_Artit-1])
}

func Filter(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	checkbox, _ := strconv.Atoi(r.Form.Get("members"))
	dateCreation, _ := strconv.Atoi(r.Form.Get("DC"))
	dateAlbum := r.Form.Get("FAD")
	countrySelection := r.Form.Get("country")

	var new_artistTab = ArtistTab
	var temp_artistTab []Artist
	for _, v := range new_artistTab {

		if (len(v.Members) == checkbox || checkbox == 0) && (v.CreationDate == dateCreation || dateCreation == 1991) && (v.FirstAlbum[6:] == dateAlbum || dateAlbum == "1991") && (ContainsCountry2(v.Concerts, countrySelection) || countrySelection == "") {
			temp_artistTab = append(temp_artistTab, v)
		}
		new_artistTab = temp_artistTab
	}
	tmpl := template.Must(template.ParseFiles("tmpl/index.html"))
	tmpl.Execute(w, ArtistStruct{Tab: new_artistTab, Country: CountryList})

}

// sort.Sort(empeo(DataRelation))

type empeo []Date

func (e empeo) Len() int {
	return len(e)
}

func (e empeo) Less(i, j int) bool {
	return e[i].Datecomp < e[j].Datecomp
}

func (e empeo) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
