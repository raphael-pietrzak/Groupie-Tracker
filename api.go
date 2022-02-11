package groupie

import (
	// "context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Variables Globales
var ArtistTab []Artist
var LocationsTab Locations
var Relation Relationnement
var Artists Artist
var variable map[string][]string
var tableau []ArtistStruct

type ArtistStruct struct {
	Tab  []Artist
	S1 Relationnement
	S2 Artist
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

type Locations struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Relationnement struct {
	Id       int                 `json:"id"`
	DatesLoc map[string][]string `json:"datesLocations"`
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

func APIRequests2(link string) {

	req, _ := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + link)

	d, _ := ioutil.ReadAll(req.Body)
	req2, _ := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + link)

	d2, _ := ioutil.ReadAll(req2.Body)

	json.Unmarshal(d, &Relation)
	json.Unmarshal(d2, &Artists)
}
