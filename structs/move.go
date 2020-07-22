package structs

type ResponseMove struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Accuracy     int    `json:"accuracy"`
	EffectChance int    `json:"effect_chance"`
	PP           int    `json:"pp"`
	Priority     int    `json:"priority"`
	Power        int    `json:"power"`
	DamageClass   DamageClass
	EffectEntries []EffectEntries
	Meta 		  Meta
	Type		  Type
	Target		  Target
	StatChanges	  StatChanges
}
type  DamageClass struct {
	Name string `json:"name"`
}
type EffectEntries struct {
	Effect   string `json:"effect"`
	Language Language
}
type  Language struct {
	Name string `json:"name"`
}
type Meta struct {
	Ailment Ailment
	Category Category
	MinHits       *int `json:"min_hits"`
	MaxHits       *int `json:"max_hits"`
	MinTurns      *int `json:"min_turns"`
	MaxTurns      *int `json:"max_turns"`
	Drain         int  `json:"drain"`
	Healing       int  `json:"healing"`
	CritRate      int  `json:"crit_rate"`
	AilmentChance int  `json:"ailment_chance"`
	FlinchChance  int  `json:"flinch_chance"`
	StatChance    int  `json:"stat_chance"`
}
type Ailment struct {
	Name string `json:"name"`
}
type Category struct {
	Name string `json:"name"`
}

type Target struct {
	Name string `json:"name"`
}
type StatChanges struct {
	Change int `json:"change"`
	Stat Stat
}


type MoveService interface {
	GetMoves() (*ResponseMove ,error)
}
type Move1 struct {
	client *Client
}
func (s *Move1) GetMoves() (*ResponseMove ,error) {

}