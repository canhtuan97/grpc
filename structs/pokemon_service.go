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

	PokemonGet PokemonService
	MoveGet    MoveService
	//Location LocationService
}

func NewClient(baseUrl string) *Client {
	httpClient := http.DefaultClient

	c := &Client{Client: httpClient, baseUrl: baseUrl}

	c.PokemonGet = &PokemonGet{client: c}
	c.MoveGet = &MoveGet{client: c}
	//c.Location = &Location{client: c}
	return c
}

//func (s *Pokemon) GetPokemon() (*ResponsePokemon, error) {
//	res, err := s.client.Client.Get(s.client.baseUrl)
//	if err != nil {
//		return nil, err
//	}
//	pokemonResponse := ResponsePokemon{}
//	responseData, err := ioutil.ReadAll(res.Body)
//	if err != nil {
//		log.Fatal(err)
//	}
//	json.Unmarshal(responseData, &pokemonResponse)
//
//	return &pokemonResponse, nil
//}

//func main() {
//
//	client := NewClient(" https://pokeapi.co/api/v2/pokemon/1")
//
//	//res , err := http.Get("https://pokeapi.co/api/v2/move/2")
//	if err != nil {
//		log.Fatal(err)
//	}
//	responseData, err := ioutil.ReadAll(res.Body)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(string(responseData))
//}
//
//
//
