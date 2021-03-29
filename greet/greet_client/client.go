package main

import (
	"context"
	"fmt"
	"gRPC/greet/greetpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
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
	//doClientStreaming(c)
	doBiDiStreaming(c)

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
func doClientStreaming(c greetpb.GreetServiceClient)  {
	fmt.Println("started ClientStreaming streaming gRPC ")

	requests:=[]* greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting:&greetpb.Greeting{
				FirstName: "zaid",

			} ,
		},
		&greetpb.LongGreetRequest{
			Greeting:&greetpb.Greeting{
				FirstName: "hamaz",

			} ,
		},
		&greetpb.LongGreetRequest{
			Greeting:&greetpb.Greeting{
				FirstName: "waqar",

			} ,
		},
		&greetpb.LongGreetRequest{
			Greeting:&greetpb.Greeting{
				FirstName: "bilal",

			} ,
		},
	}
	stream,err:=c.LongGreet(context.Background())
	if err!=nil{
		fmt.Println("ERROR sending streaming request ",err)
	}
	for index, element := range requests {
		fmt.Println("senfing message number ",index)
		stream.Send(element)
		time.Sleep(1000 *time.Millisecond)


	}

	res,err:=stream.CloseAndRecv()
	if err!=nil {
		fmt.Println("error reading message from server ",err)
	}
	fmt.Println(res)



}
func doBiDiStreaming(c greetpb.GreetServiceClient)  {
	fmt.Println("do bidirectional Invoked")

	requests:=[]* greetpb.GreetEveryoneRequest{
		&greetpb.GreetEveryoneRequest{
			Greeting:&greetpb.Greeting{
				FirstName: "zaid",

			} ,
		},
		&greetpb.GreetEveryoneRequest{
			Greeting:&greetpb.Greeting{
				FirstName: "hamaz",

			} ,
		},
		&greetpb.GreetEveryoneRequest{
			Greeting:&greetpb.Greeting{
				FirstName: "waqar",

			} ,
		},
		&greetpb.GreetEveryoneRequest{
			Greeting:&greetpb.Greeting{
				FirstName: "bilal",

			} ,
		},
	}

	res,err:=c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("ERROr from server ",err)
		return
	}
	waitc:=make(chan struct{})
	go func() {
		for _,d:= range requests{
			fmt.Println("Sending message ",d)
			res.Send(d)
		}
		res.CloseSend()
	}()
	go func() {

		for  {
			response , err:=res.Recv()
			if err ==io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("ERRRO recoiving from server stream ")
				break
			}
			fmt.Println("recived from server ",response.GetResult())
		}

		close(waitc)

	}()


	//block until everything is done
	<-waitc

}