package main

import (
	"Golang-gRPC/greet/greetpb"
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
	c := greetpb.NewGreetServiceClient(conn)
	//fmt.Println("created client ",c)
	doUnary(c)

}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("started unary gRPC ")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "zaid ",
			LastName:  "saeed",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		fmt.Println("errror sending request")
	}
	fmt.Println("SUCCESS SENDING REQUEST", res)

}
