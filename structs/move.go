package structs

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type ResponseMove struct {
	ID                 int                 `json:"id"`
	Name               string              `json:"name"`
	Accuracy           int                 `json:"accuracy"`
	EffectChance       int                 `json:"effect_chance"`
	PP                 int                 `json:"pp"`
	Priority           int                 `json:"priority"`
	Power              int                 `json:"power"`
	ContestCombos      ContestCombos       `json:"contest_combos"`
	ContestType        ContestType         `json:"contest_type"`
	ContestEffect      ContestEffect       `json:"contest_effect"`
	DamageClass        DamageClass         `json:"damage_class"`
	EffectEntries      EffectEntries       `json:"effect_entries"`
	EffectChanges      EffectChanges       `json:"effect_changes"`
	Generation         Generation          `json:"generation"`
	Meta               Meta                `json:"meta"`
	Names              []Names             `json:"names"`
	PastValues         []PastValues        `json:"past_values"`
	StatChanges        []StatChanges       `json:"stat_changes"`
	SuperContestEffect SuperContestEffect  `json:"super_contest_effect"`
	Target             Target              `json:"target"`
	Type               Type                `json:"type"`
	FlavorTextEntries  []FlavorTextEntries `json:"flavor_text_entries"`
}

type PastValues struct {

}

type StatChanges struct {

}
type ContestCombos struct {
	Normal Normal `json:"normal"`
	Super  Super  `json:"super"`
}

type Normal struct {
	UserBefore UserBefore
	UserAfter string
}

type UserBefore struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

type Super struct {
	UseBefore string `json:"use_before"`
	UseAfter string `json:"use_after"`
}

type ContestType struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type ContestEffect struct {
	Url string `json:"url"`
}

type DamageClass struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type EffectEntries struct {
	Effect      string `json:"effect"`
	ShortEffect string `json:"short_effect"`
	Language    Language
}
type Language struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type EffectChanges struct {

}

type Generation struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
type Meta struct {
	Ailment       Ailment  `json:"ailment"`
	Category      Category `json:"category"`
	MinHits       int      `json:"min_hits"`
	MaxHits       int      `json:"max_hits"`
	MinTurns      int      `json:"min_turns"`
	MaxTurns      int      `json:"max_turns"`
	Drain         int      `json:"drain"`
	Healing       int      `json:"healing"`
	CritRate      int      `json:"crit_rate"`
	AilmentChance int      `json:"ailment_chance"`
	FlinchChance  int      `json:"flinch_chance"`
	StatChance    int      `json:"stat_chance"`
}

type Ailment struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Category struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Names struct {
	Name string `json:"name"`
	Language Language `json:"language"`
}

type SuperContestEffect struct {
	Url string `json:"url"`
}

type Target struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type FlavorTextEntries struct {
	FlavorText   string       `json:"flavor_text"`
	Language     Language     `json:"\"`
	VersionGroup VersionGroup `json:"version_group"`
}

type MoveService  interface {
	GetMoves() (*ResponseMove , error)
}
type MoveGet struct {
	client *Client
}


func (s *MoveGet) GetMoves() (*ResponseMove ,error) {
	res, err := s.client.Client.Get(s.client.baseUrl)
	if err != nil {
		return nil, err
	}
	moveResponse := ResponseMove{}
	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(responseData, &moveResponse)

	return &moveResponse, nil
}