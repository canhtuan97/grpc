package structs

import (
	"fmt"
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

	Pokemon  PokemonService
	//Move1    MoveService
	//Location LocationService
}

func NewClient(baseUrl string) *Client {
	httpClient := http.DefaultClient

	c := &Client{Client: httpClient, baseUrl: baseUrl}

	c.Pokemon = &Pokemon{client: c}
	//c.Move1 = &Move1{client: c}
	//c.Location = &Location{client: c}
	return c
}


func main() {

	client := NewClient(" https://pokeapi.co/api/v2/pokemon/1")

	//res , err := http.Get("https://pokeapi.co/api/v2/move/2")
	if err != nil {
		log.Fatal(err)
	}
	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
}



