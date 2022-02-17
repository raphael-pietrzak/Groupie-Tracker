package groupie

import (
	// "context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Variables Globales
var Relation Relationnement
var Artists Artist
var ArtistTab []Artist
var LocationTab []Location
var L Loc
var variable map[string][]string
var tableau []ArtistStruct
var Countries []string

type ArtistStruct struct {
	Tab []Artist
	S1  []Date
}
type Loc struct {
	Index  []Location `json:"index"`
}

type Location struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	Dates        string   `json:"dates"`
	Relations    string   `json:"relations"`
}

type Relationnement struct {
	Id       int                 `json:"id"`
	DatesLoc map[string][]string `json:"datesLocations"`
}

type Date struct {
	Day     string
	Month   string
	Year    string
	City    string
	Country string
}

func APIRequests() {

	req, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	if err != nil {
		fmt.Println(err)
	}

	d, err2 := ioutil.ReadAll(req.Body)

	if err2 != nil {
		fmt.Println(err2)
	}

	json.Unmarshal(d, &ArtistTab)
}

func APIRequestsLoc() {

	req, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")

	if err != nil {
		fmt.Println(err)
	}

	d3, err2 := ioutil.ReadAll(req.Body)

	if err2 != nil {
		fmt.Println(err2)
	}

	json.Unmarshal(d3, &L)
	LocationTab = L.Index
}

func APIRequests2(link string) {

	req, _ := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + link)

	d1, _ := ioutil.ReadAll(req.Body)
	req2, _ := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + link)

	d2, _ := ioutil.ReadAll(req2.Body)

	json.Unmarshal(d1, &Relation)
	json.Unmarshal(d2, &Artists)

}
