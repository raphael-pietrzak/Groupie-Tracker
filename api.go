package groupie

import (
<<<<<<< HEAD
	//"enconding/json"
=======
>>>>>>> 1c15842413d449bcedea0f07b768a934386d0ff3
	"fmt"
	"io/ioutil"
	"net/http"
)

type ArtistStruct struct {
	tab []Artist
}

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
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

func APIRequests() {

	req, err := http.Get("https://groupietrackers.herokuapp.com/api")

	if err != nil {
		fmt.Println(err)
	}

	d, err2 := ioutil.ReadAll(req.Body)

	if err2 != nil {
		fmt.Println(err2)
	}
}
