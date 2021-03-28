package main

import (
	"context"
	"fmt"
	"gRPC/greet/greetpb"
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
	c:=greetpb.NewGreetServiceClient(conn)
	//fmt.Println("created client ",c)
	doServerStreaming(c)

}

func doUnary(c greetpb.GreetServiceClient)  {
fmt.Println("started unary gRPC ")
	req:=&greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "zaid ",
			LastName:  "saeed",
		},

	}
	res, err := c.Greet(context.Background(), req)
	if err!=nil {
		fmt.Println("errror sending request")
	}
	fmt.Println("SUCCESS SENDING REQUEST",res)
}


func doServerStreaming(c greetpb.GreetServiceClient)  {
	fmt.Println("started ServerStreaming streaming gRPC ")
	req:=&greetpb.GreetManytimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "zaid ",
			LastName:  "saeed",
		},

	}
	StreamRes, err := c.GreetManytimes(context.Background(), req)
	if err!=nil {
		fmt.Println("errror sending GreetManytimes")
	}
	for{
		msg, err :=StreamRes.Recv()
		if err!=nil{
			fmt.Println("errror while reading  messsage GreetManytimes")
			break

		}
		fmt.Println("SUCCESS SENDING REQUEST",msg.GetResult())
	}

}
