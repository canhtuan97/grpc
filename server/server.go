package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/canhtuan97/grpc/proto"
	"github.com/canhtuan97/grpc/structs"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (*server) GetPokemon(ctx context.Context, req *app_serverpb.MoveRequest) (*app_serverpb.MoveResponse, error) {
	panic("implement me")
}

//func (s server) GetPokemon(ctx context.Context, request *app_serverpb.MoveRequest) (*app_serverpb.MoveResponse, error) {
//	panic("implement me")
//}

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

	lis ,err := net.Listen("tpc","0.0.0.0:50069")
	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}

	s := grpc.NewServer()

	app_serverpb.RegisterPokemonServiceServer(s , &server{})

	fmt.Println("Server running ...")

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("err while serve %v", err)
	}
}