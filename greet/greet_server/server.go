package main

import (
	"Golang-gRPC/greet/greetpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Println("great function was invoked")
	firstName := req.GetGreeting().FirstName
	result := "He llo " + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil

}

func main() {
	fmt.Println("heyy")
	lis, err := net.Listen("tcp", "0.0.0.0:5855")
	if err != nil {
		log.Fatal("Failed,", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed ", err)
	}

}
