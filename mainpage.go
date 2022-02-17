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
	testvar := []string{}
	for num, _ := range ArtistTab {
		APIRequests2(strconv.Itoa(num))
		for k, _ := range Relation.DatesLoc{
			split := strings.Split(k,"-")
			country := strings.Title((strings.Replace(split[1],"_"," ",-1))) 
			if containsCountry(testvar,country) == false {
				testvar = append(testvar, country)  
			}
		}
	}
	sort.Strings(testvar)
	fmt.Println(testvar)

	new := ArtistStruct{Tab: ArtistTab}
	tmpl.Execute(w, new)
}

func containsCountry(testvar []string, str string) bool {
	for _, v := range testvar {
		if v == str {
			return true
		}
	}
	return false
}

func Artiste(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	link_loc := r.Form.Get("w")

	tmpl := template.Must(template.ParseFiles("tmpl/artist.html"))

	APIRequests2(link_loc)
	DataRelation :=  []Date{}
	var DRelation Relationnement
	month := []string{"JAN","FEB","MAR","APR","MAY","JUN","JUL","AUG","SEP","OCT","NOV","DEC"}
	for k, v := range Relation.DatesLoc {
		a := Date{}
		split := strings.Split(k, "-")
		city := strings.Title((strings.Replace(split[0], "_", " ", -1)))
		country := strings.Title((strings.Replace(split[1], "_", " ", -1)))
		for _,i := range v{
			splitdate := strings.Split(i, "-")
			m,_ := strconv.Atoi(splitdate[1])
			a.Day = splitdate[0]
			a.Month = month[m-1]
			a.Year = splitdate[2]
			a.City = city
			a.Country = country 
			DataRelation = append(DataRelation, a)
		}
	}
	DRelation.Id,_ = strconv.Atoi(link_loc)
	tmpl.Execute(w, ArtistStruct{ S1: DataRelation, Tab: []Artist{Artists}})
}

