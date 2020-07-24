package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/canhtuan97/grpc/proto_demo"

	"google.golang.org/grpc"
)

type server struct{}

func (server) Echo(ctx context.Context, msg *example.StringMessage) (*example.StringMessage, error) {
	log.Printf("receive message %s\n", msg.GetValue())
	return msg, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50080")
	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}

	s := grpc.NewServer()

	example.RegisterYourServiceServer(s, &server{})

	fmt.Println("demo gateway service is running...")
	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("err while serve %v", err)
	}
}
