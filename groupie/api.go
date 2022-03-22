package groupie

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var Artists Artist
var CountryList []string
var ArtistTab []Artist
var DatesLocations DatesLoc

type ArtistStruct struct {
	Tab     []Artist
	Country []string
}

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Concerts     []Location
}

type DatesLoc struct {
	DatesLoc map[string][]string `json:"datesLocations"`
}

type Location struct {
	City    string
	Country string
	Date    Date
}

type Date struct {
	Day      string
	Month    string
	Year     string
	Datecomp int
}

func APIRequests() {

	req, _ := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	artists, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(artists, &ArtistTab)

	month := []string{"JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"}

	for k := range ArtistTab {
		n := strconv.Itoa(k + 1)
		DatesLocations := DatesLoc{}
		req2, _ := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + n)
		locations, _ := ioutil.ReadAll(req2.Body)
		json.Unmarshal(locations, &DatesLocations)
		for key, value := range DatesLocations.DatesLoc {
			splitloc := strings.Split(key, "-")
			city := strings.Title((strings.Replace(splitloc[0], "_", " ", -1)))
			country := strings.Title((strings.Replace(splitloc[1], "_", " ", -1)))
			for _, i := range value {
				splitdate := strings.Split(i, "-")
				m, _ := strconv.Atoi(splitdate[1])
				datecomp, _ := strconv.Atoi(splitdate[2] + splitdate[1] + splitdate[0])
				ArtistTab[k].Concerts = append(ArtistTab[k].Concerts, Location{City: city, Country: country, Date: Date{Day: splitdate[0], Month: month[m-1], Year: splitdate[2], Datecomp: datecomp}})
			}
			if ContainsCountry(CountryList, country) {
				CountryList = append(CountryList, country)
			}
		}

	}
}

func ContainsCountry(testvar []string, str string) bool {
	for _, v := range testvar {
		if v == str {
			return false
		}
	}
	return true
}

func NumMembers(testvar []int, i int) bool {

	for _, v := range testvar {
		if v == i {
			return true
		}
	}
	return false
}
