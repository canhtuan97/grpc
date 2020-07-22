package structs
import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type ResponsePokemon struct {
	Id                     int                      `json:"id"`
	Name                   string                   `json:"name"`
	BaseExperience         int                      `json:"base_experience"`
	Height                 int                      `json:"height"`
	IsDefault              int                      `json:"is_default"`
	Order                  int                      `json:"order"`
	Weight                 int                      `json:"weight"`
	Abilities              []Abilities              `json:"abilities"`
	Forms                  []Forms                  `json:"forms"`
	GameIndices            []GameIndices            `json:"game_indices"`
	HeldItems              []HeldItems              `json:"held_items"`
	LocationAreaEncounters []LocationAreaEncounters `json:"location_area_encounters"`
	Moves                  []Moves                  `json:"moves"`
	Species                Species                  `json:"species"`
	Sprites                Sprites                  `json:"sprites"`
	Stats                  []Stats                  `json:"stats"`
	Types                  []Types                  `json:"types"`
}

type Abilities struct {
	IsHidden bool    `json:"is_hidden"`
	Slot     int     `json:"slot"`
	Ability  Ability `json:"ability"`
}

type Ability struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Forms struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type GameIndices struct {
	GameIndex int     `json:"game_index"`
	Version   Version `json:"version"`
}

type Version struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type HeldItems struct {
	Item           Item           `json:"item"`
	VersionDetails VersionDetails `json:"version_details"`
}

type Item struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type VersionDetails struct {
	Rarity           int                `json:"rarity"`
	Version          Version            `json:"version"`
	MaxChance        int                `json:"max_chance"`
	EncounterDetails []EncounterDetails `json:"encounter_details"`
}

type EncounterDetails struct {
	MinLevel        int               `json:"min_level"`
	MaxLevel        int               `json:"max_level"`
	ConditionValues []ConditionValues `json:"condition_values"`
	Chance          int               `json:"chance"`
	Method          Method            `json:"method"`
}

type Method struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type ConditionValues struct {
	Name string `json:"name"`
	url  string `json:"url"`
}

type LocationAreaEncounters struct {
	LocationArea   LocationArea     `json:"location_area"`
	VersionDetails []VersionDetails `json:"version_details"`
}

type LocationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Moves struct {
	Move Move `json:"move"`
	VersionGroupDetails []VersionGroupDetails `json:"version_group_details"`
}

type Move struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type VersionGroupDetails struct {
	LevelLearnedAt  int             `json:"level_learned_at"`
	VersionGroup    VersionGroup    `json:"version_group"`
	MoveLearnMethod MoveLearnMethod `json:"move_learn_method"`
}

type VersionGroup struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
type MoveLearnMethod struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Species struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Sprites struct {
	BackFemale       string `json:"back_female"`
	BackShinyFemale  string `json:"back_shiny_female"`
	BackDefault      string `json:"back_default"`
	FrontFemale      string `json:"front_female"`
	FrontShinyFemale string `json:"front_shiny_female"`
	BackShiny        string `json:"back_shiny"`
	FrontDefault     string `json:"front_default"`
	FrontShiny       string `json:"front_shiny"`
}

type Stats struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     Stat `json:"stat"`
}
type Stat struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

type Types struct {
	Slot int  `json:"slot"`
	Type Type `json:"type"`
}

type Type struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonService interface {
	GetPokemon() (*ResponsePokemon, error)
}

type Pokemon struct {
	client *Client
}

func (s *Pokemon) GetPokemon() (*ResponsePokemon, error) {
	res, err := s.client.Client.Get(s.client.baseUrl)
	if err != nil {
		return nil, err
	}
	pokemonResponse := ResponsePokemon{}
	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(responseData, &pokemonResponse)

	return &pokemonResponse, nil
}