package main

import (
	"fmt"
	"context"
	"log"
	"net"
	
	"google.golang.org/grpc"
	"github.com/canhtuan97/grpc/calculatorpb"
	"github.com/mtslzr/pokeapi-go"
	
)
type server struct{}

func Move(id string) (result structs.Move, err error) {
	err = do(fmt.Sprintf("move/%s", id), &result)
	return result, err
}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	log.Println("sum called...")
	resp := &calculatorpb.SumResponse{
		Result: req.GetNum1() + req.GetNum2(),
	}

	return resp, nil
}


func main()  {
	lis,err := net.Listen("tcp","0.0.0.0:50069")
	if err != nil {
		log.Fatalf("err while create listen  %v",err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})
	
	fmt.Println("calculator is running...")
	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("err while serve %v",err)
	}
}