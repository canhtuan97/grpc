package main

import (
	"log"
	"context"

	"google.golang.org/grpc"
	"github.com/canhtuan97/grpc/calculatorpb"
)

func callSum(c calculatorpb.CalculatorServiceClient) {
	log.Println("calling sum api")
	resp, err := c.Sum(context.Background(), &calculatorpb.SumRequest{
		Num1: 5,
		Num2: 6,
	})

	if err != nil {
		log.Fatalf("call sum api err %v", err)
	}

	log.Printf("sum api response %v\n", resp.GetResult())
}

func main(){
	cc,err := grpc.Dial("localhoast:50069",grpc.WithInsecure())

	if err !=nil {
		log.Fatalf("err while dial %v" ,err)
	}

	// dong ket noi
	defer cc.Close()

	client := calculatorpb.NewCalculatorServiceClient(cc)
	callSum(client)
	log.Printf("service client %f" ,client)
}