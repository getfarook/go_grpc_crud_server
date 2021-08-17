package main

import (
	"context"
	"fmt"
	"log"

	"github.com/simplesteph/grpc-go-course/calculator/calcpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm the client ")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to server %v", err)

	}

	defer cc.Close()

	c := calcpb.NewCalcServiceClient(cc)

	doUnary(c)

}

func doUnary(c calcpb.CalcServiceClient) {
	fmt.Println("Starting to do Unary RPC")

	req := &calcpb.CalcRequest{
		Inputs: &calcpb.Inputs{
			X: 20,
			Y: 30,
		},
	}

	res, err := c.Calc(context.Background(), req)
	if err != nil {
		log.Fatalf("error while callin Calc RPC: %v", err)
	}
	log.Printf("Response from Calc: %v", res.Result)

}
