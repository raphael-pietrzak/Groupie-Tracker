package groupie

import (
	// "fmt"

	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("tmpl/index.html"))

	APIRequests()

	new := ArtistStruct{Tab: ArtistTab}

	// fmt.Println(new)
	tmpl.Execute(w, new)
}

func Filter(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	checkbox, _ := strconv.Atoi(r.Form.Get("members"))
	dateCreation, _ := strconv.Atoi(r.Form.Get("slider_DC"))
	dateAlbum := r.Form.Get("slider_FAD")
	fmt.Println(checkbox)
	fmt.Println(dateCreation)
	fmt.Println(dateAlbum)

	var new_artistTab = ArtistTab
	if checkbox != 0 {
		var temp_artistTab []Artist
		for _, v := range new_artistTab {
			if len(v.Members) == checkbox {
				temp_artistTab = append(temp_artistTab, v)
			}

		new_artistTab = temp_artistTab
		}
		// split := strings.Split(v.FirstAlbum, "-")
		// aza := split[2]
		// if checkbox != 0 {
		// 	if dateCreation != 1987 {
		// 		if len(v.Members) == checkbox && v.CreationDate >= dateCreation {
		// 			new_artistTab = append(new_artistTab, v)
		// 		} else {
		// 			if v.FirstAlbum == dateAlbum {
		// 				new_artistTab = append(new_artistTab, v)
		// 			}
		// 		}
		// 	} else {
		// 		if len(v.Members) >= checkbox {
		// 			new_artistTab = append(new_artistTab, v)
		// 		}
		// 	}
		// } else {
		// 	if v.CreationDate >= dateCreation {
		// 		new_artistTab = append(new_artistTab, v)
		// 	}
		// }
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
	fmt.Println(DataRelation)
	
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
			a.City = city
			a.Country = country
			DataRelation = append(DataRelation, a)
		}
	}
	DRelation.Id, _ = strconv.Atoi(link_loc)
	tmpl.Execute(w, ArtistStruct{S1: DataRelation, Tab: []Artist{Artists}})
}
