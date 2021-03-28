package main

import (
	"context"
	"fmt"
	"gRPC/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
)

func main()  {
	fmt.Println("client is runnnig")
	conn,err:=grpc.Dial("localhost:5855",grpc.WithInsecure())
	if err!=nil{
		log.Fatalf("could not connect ",err)
	}
	defer conn.Close()
	c:=calculatorpb.NewCalculatorServiceClient(conn)
	//fmt.Println("created client ",c)
	doServerStreaming(c)

}

func doUnary(c calculatorpb.CalculatorServiceClient)  {
	fmt.Println("started unary gRPC ")
	req:=&calculatorpb.SumRequest{
		FirstNumber:  2,
		SecondNumber: 3,
	}
	res, err := c.Sum(context.Background(), req)
	if err!=nil {
		fmt.Println("errror sending request")
	}
	fmt.Println("SUCCESS SENDING REQUEST",res)



}

func doServerStreaming(c calculatorpb.CalculatorServiceClient)  {
	fmt.Println("started Server streaming  gRPC ")
	req:=&calculatorpb.PrimeNumberDecompositionRequest{
		Number: 132135163444,
	}
	res, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err!=nil {
		fmt.Println("errror sending request")
	}
	fmt.Println("SUCCESS SENDING REQUEST",res)

	for{
		msg , err :=res.Recv()
		if err!=nil{
			println("errro getting msg")
			break
		}

		fmt.Println("getting message ",		msg.GetPrimefactor())


	}




}