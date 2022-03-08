package main

import (
	"Golang-gRPC/calculator/calculatorpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("client is runnnig")
	conn, err := grpc.Dial("localhost:5855", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect ", err)
	}
	defer conn.Close()
	c := calculatorpb.NewCalculatorServiceClient(conn)
	//fmt.Println("created client ",c)
	doUnary(c)

}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("started unary gRPC ")
	req := &calculatorpb.SumRequest{
		FirstNumber:  2,
		SecondNumber: 3,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		fmt.Println("errror sending request")
	}
	fmt.Println("SUCCESS SENDING REQUEST", res)

}
