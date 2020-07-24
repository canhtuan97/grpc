package structs

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//type PokemonService interface {
//	GetPokemons()
//	GetMoves()
//	GetLocations()
//}

//type

type PokemonServiceInstance struct {
	Client *Client

}

type Client struct {
	Client  *http.Client
	baseUrl string

	PokemonGet PokemonService
	MoveGet    MoveService
	//Location LocationService
}

func NewClient(baseUrl string) *Client {
	httpClient := http.DefaultClient

	c := &Client{Client: httpClient, baseUrl: baseUrl}

	c.PokemonGet = &PokemonGet{client: c}
	c.MoveGet = &MoveGet{client: c}

	return c
}

//MapUnmarshal
func MapUnmarshal(res *http.Response,resouce interface{})   error  {
	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	dataConvert := json.Unmarshal(responseData, &resouce)
	return dataConvert

}

