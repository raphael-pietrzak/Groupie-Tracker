package groupie

import (
	// "fmt"

	"fmt"
	"html/template"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("tmpl/index.html"))

	APIRequests()

	new := ArtistStruct{Countries: Countries, Tab: ArtistTab}

	// fmt.Println(new)
	tmpl.Execute(w, new)
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

	fmt.Println("checkbox")
	fmt.Println(checkbox)
	fmt.Println("dateCreation")
	fmt.Println(dateCreation)
	fmt.Println("dateAlbum")
	fmt.Println(dateAlbum)
	fmt.Println("countrySelection")
	fmt.Println(countrySelection)

	var new_artistTab = ArtistTab
	var temp_artistTab []Artist
	for _, v := range new_artistTab {
		splitdate := strings.Split(v.FirstAlbum, "-")
		year := splitdate[2]
		split := strings.Split(v.Locations, "-")
		fmt.Println(split)
		country := strings.Title((strings.Replace(split[1], "_", " ", -1)))
		if (len(v.Members) == checkbox || checkbox == 0) && (v.CreationDate == dateCreation || dateCreation == 1991 ) && (year == dateAlbum || dateAlbum == "1991" ) && (country == countrySelection || countrySelection == ""){
			temp_artistTab = append(temp_artistTab, v)
		}
	}

	if dateCreation < 10000 {
		for _, i := range new_artistTab {
			if i.CreationDate <= dateCreation {
				temp_artistTab = append(temp_artistTab, i)
			}
		}
		new_artistTab = temp_artistTab
	}

	tmpl := template.Must(template.ParseFiles("tmpl/index.html"))
	tmpl.Execute(w, ArtistStruct{Tab: new_artistTab})

}

func ContainsCountry(testvar []string, str string) bool {
	for _, v := range testvar {
		if v == str {
			return false
		}
	}
	return true
}

// func containsCountry(testvar []string, str string) bool {
// 	for _, v := range testvar {
// 		if v == str {
// 			return true
// 		}
// 	}
// 	return false
// }

func Artiste(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	link_loc := r.Form.Get("w")

	tmpl := template.Must(template.ParseFiles("tmpl/artist.html"))

	APIRequests2(link_loc)
	DataRelation := []Date{}
	var DRelation Relationnement
	month := []string{"JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"}
	for k, v := range Relation.DatesLoc {
		a := Date{}
		split := strings.Split(k, "-")
		city := strings.Title((strings.Replace(split[0], "_", " ", -1)))
		country := strings.Title((strings.Replace(split[1], "_", " ", -1)))
		for _, i := range v {
			splitdate := strings.Split(i, "-")
			m, _ := strconv.Atoi(splitdate[1])
			a.Day = splitdate[0]
			a.Month = month[m-1]
			a.Year = splitdate[2]
			a.Datecomp, _ = strconv.Atoi(splitdate[2]+splitdate[1]+splitdate[0])
			a.City = city
			a.Country = country
			DataRelation = append(DataRelation, a)
		}
	}
	DRelation.Id, _ = strconv.Atoi(link_loc)

	
	
	sort.Sort(empeo(DataRelation))
	tmpl.Execute(w, ArtistStruct{S1: DataRelation, Tab: []Artist{Artists}})
}

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
