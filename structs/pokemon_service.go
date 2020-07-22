package structs

import (
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
	Move1    MoveService
	Location LocationService
}

func NewClient(baseUrl string) *Client {
	httpClient := http.DefaultClient

	c := &Client{Client: httpClient, baseUrl: baseUrl}

	c.Pokemon = &Pokemon{client: c}
	c.Move1 = &Move1{client: c}
	c.Location = &Location{client: c}
	return c
}


func main() {

}



