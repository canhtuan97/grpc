package main

import (
	"encoding/json"
	"fmt"
	"github.com/canhtuan97/grpc/structs"
	"log"
)

func Test()  {
	fmt.Printf("Starting server at port 8080\n")

	client := structs.NewClient("https://pokeapi.co/api/v2/move/1")
	data, err := client.MoveGet.GetMoves()
	if err != nil {
		log.Fatal(err)
	}
	strDat, _ := json.Marshal(data)
	fmt.Println("string data: ", string(strDat))
}
func main() {



}