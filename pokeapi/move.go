package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"


)

type Move struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Accuracy     int    `json:"accuracy"`
	EffectChance int    `json:"effect_chance"`
	PP           int    `json:"pp"`
	Priority     int    `json:"priority"`
	Power        int    `json:"power"`
	DamageClass  struct {
		Name string `json:"name"`
	} `json:"damage_class"`
	EffectEntries []struct {
		Effect   string `json:"effect"`
		Language struct {
			Name string `json:"name"`
		} `json:"language"`
	} `json:"effect_entires"`
	Meta struct {
		Ailment struct {
			Name string `json:"name"`
		} `json:"ailment"`
		Category struct {
			Name string `json:"name"`
		}
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
	} `json:"meta"`
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
	Target struct {
		Name string `json:"name"`
	} `json:"target"`
	StatChanges []struct {
		Change int `json:"change"`
		Stat   struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stat_changes"`
}



func main() {
	response, err := http.Get("https://pokeapi.co/api/v2/move/2")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var move Move
	json.Unmarshal(responseData, &move)
	//fmt.Println(string(responseData))
	fmt.Printf("id: %s, Accuracy: %s,PP : %s", move.ID, move.Accuracy,move.PP)

}