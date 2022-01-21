package groupie

import (
	//"enconding/json"
	"fmt"
	"io/ioutil"
	"io/outil"
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
