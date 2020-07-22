package main

import (
	"encoding/json"
	"fmt"

	//"../structs"
	"github.com/canhtuan97/grpc/structs"
	"log"

)




func main()  {

	client := structs.NewClient("https://pokeapi.co/api/v2/pokemon/1")
	data, err := client.Pokemon.GetPokemon()
	if err != nil {
		log.Fatal(err)
	}
	strDat, _ := json.Marshal(data)
	fmt.Println("string data: ", string(strDat))

}