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
var Relation Relations
var variable map[string][]string

type ArtistStruct struct {
	Tab  []Artist
	Tab2 []string
	// Tab3 Locations
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
}

type Locations struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Relations struct {
	Index []Relationnement `json:"index`
}

type Relationnement struct {
	Id       int            `json:"id"`
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

	req, err := http.Get(link)

	if err != nil {
		fmt.Println(err)
	}

	d, err2 := ioutil.ReadAll(req.Body)

	if err2 != nil {
		fmt.Println(err2)
	}
	json.Unmarshal(d, &LocationsTab)


	// Relier ça... 

	fmt.Println(LocationsTab.Id)
	fmt.Println(LocationsTab.Locations)

	req2, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")


	if err != nil {
		fmt.Println(err)
	}

	d2, err2 := ioutil.ReadAll(req2.Body)

	if err2 != nil {
		fmt.Println(err2)
	}
	json.Unmarshal(d2, &Relation)


	// Avec ça !!
	fmt.Println(Relation)
	for _, v := range Relation.Index {
		if v.Id == LocationsTab.Id{
			variable = v.DatesLoc
		}
	}
}



