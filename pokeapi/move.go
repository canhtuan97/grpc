package main

import (

	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"


)


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
	fmt.Println(string(responseData))
	//var move structs.Move
	//json.Unmarshal(responseData, &move)
	////fmt.Println(string(responseData))
	//fmt.Printf("id: %s, Accuracy: %s,PP : %s", move.ID, move.Accuracy,move.PP)

}