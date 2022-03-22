package groupie

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func MainPage(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("groupie/tmpl/index.html"))

	new := ArtistStruct{Tab: ArtistTab, Country: CountryList}
	tmpl.Execute(w, new)
}

func Artiste(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("groupie/tmpl/artist.html"))

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
	checkbox := []int{}
	for i := 0; i < 8; i++ {
		
		a,_ := strconv.Atoi(r.Form.Get("members"+strconv.Itoa(i)))
		if a != 0{
			checkbox = append(checkbox, a)
		}
	}
	dateCreation, _ := strconv.Atoi(r.Form.Get("DC"))
	dateAlbum := r.Form.Get("FAD")
	countrySelection := r.Form.Get("country")

	var new_artistTab = ArtistTab
	var temp_artistTab []Artist
	for _, v := range new_artistTab {

		if (NumMembers(checkbox,len(v.Members)) || len(checkbox)==0) && (v.CreationDate == dateCreation || dateCreation == 1974) && (v.FirstAlbum[6:] == dateAlbum || dateAlbum == "1974") && ( ContainsCountrymusk(v.Concerts, countrySelection) || countrySelection == "" ) {
			temp_artistTab = append(temp_artistTab, v)
		}
		new_artistTab = temp_artistTab
	}

	tmpl := template.Must(template.ParseFiles("groupie/tmpl/index.html"))
	tmpl.Execute(w, ArtistStruct{Tab: new_artistTab, Country: CountryList})

}

func ContainsCountrymusk(testvar []Location, str string) bool {
	fmt.Println("zjdfhiouazehf")
	for _, v := range testvar {
		if v.Country == str {
			fmt.Println(v.Country, str)
			return true
		}
	}
	return false
}