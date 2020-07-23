package main

import (
	"context"
	"log"

	"github.com/canhtuan97/grpc/proto"
	"google.golang.org/grpc"
)


func callGetPokemon(p app_serverpb.PokemonServiceClient) {
	log.Println("Get pokemon is running...")
	_, err := p.GetPokemon(context.Background(), &app_serverpb.PokemonRequest{
		Url: "https://pokeapi.co/api/v2/move/1",
	})
	if err != nil {
		log.Fatalf("call Getpokemon api err %v", err)
	}

}

func main(){
	cc,err := grpc.Dial("127.0.0.1:50069",grpc.WithInsecure())

	if err !=nil {
		log.Fatalf("err while dial %v" ,err)
	}

	// dong ket noi
	defer cc.Close()

	//client := app_serverpb.NewCalculatorServiceClient(cc)
	client := app_serverpb.NewPokemonServiceClient(cc)
	callGetPokemon(client)
}