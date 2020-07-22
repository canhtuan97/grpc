package structs
import (
	"encoding/json"
	"io/ioutil"
	"log"
)
type ResponsePokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	BaseExperience int    `json:"base_experience"`
	Types 		   []Types
	Sprites 	   Sprites
	Stats		   []Stats
	Moves 		   []Moves

}

type Types struct {
	Slot int `json:"slot"`
	Type types
}

type types struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Sprites struct {
	FrontDefault string `json:"front_default"`
}

type Stats struct {
	BaseStat int `json:"base_stat"`
	Effort   int `json:"effort"`
	Stat 	 stat

}

type  stat struct {
	Name string `json:"name"`
}
type  Moves struct {
	Move Move
	VersionGroupDetails []VersionGroupDetails
}
type Move struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type VersionGroupDetails struct {
	LevelLearnedAt  int `json:"level_learned_at"`
	MoveLearnMethod MoveLearnMethod
}
type MoveLearnMethod struct {
	Name string `json:"name"`
}


type PokeRequest struct {
	Url string
}

type PokemonService interface {
	GetPokemon() (*ResponsePokemon, error)
}

type Pokemon struct {
	client *Client
}

func (s *Pokemon) GetPokemon() (*ResponsePokemon , error) {
	res, err := s.client.Client.Get(s.client.baseUrl)
	if err != nil {
		return nil,err
	}
	pokemonResponse := &ResponsePokemon{}
	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(responseData, &pokemonResponse)
	////fmt.Println(string(responseData))
	return pokemonResponse, nil
}