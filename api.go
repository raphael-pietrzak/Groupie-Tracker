package groupie

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ArtistStruct struct {
	Tab  []Artist
}

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
}

type Locations struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Date struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}
type DatesLocations struct {
	Location string   `json:"locations"`
	Date     []string `json:"dates"`
}

type Relations struct {
	Id             int            `json:"id"`
	Dateslocations DatesLocations `json:"dateslocations"`
}

var ArtistTab []Artist
var LocationTab []Artist

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

func APIRequests2() {
	req, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	if err != nil {
		fmt.Println(err)
	}

	d, err2 := ioutil.ReadAll(req.Body)

	if err2 != nil {
		fmt.Println(err2)
	}

	json.Unmarshal(d, &LocationTab)
	
}
