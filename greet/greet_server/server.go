package main


import (
	"context"
	"fmt"
	"gRPC/greet/greetpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

type server struct {
}
func (*server)Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error){
	fmt.Println("great function was invoked")
	firstName:= req.GetGreeting().FirstName
	result:= "He llo "+firstName
	res:=&greetpb.GreetResponse{
		Result: result,
	}
	return res,nil

}
func (*server)GreetManytimes(req *greetpb.GreetManytimesRequest,stream greetpb.GreetService_GreetManytimesServer)  error{
	firsNam:= req.GetGreeting().FirstName

	for i:=0 ;i <10; i++ {
		result:="Hello my name is "+firsNam + strconv.Itoa(i)
		res:=&greetpb.GreetManytimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Microsecond)
	}
	return nil

}
func (*server)LongGreet(stream greetpb.GreetService_LongGreetServer) error{
fmt.Println("LongGreet was invoked with streaming client request  ",stream)
	result:=""
	for{
		req, err:=stream.Recv()
		if err == io.EOF {//we have finshed reading client stream
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})

		}
		if err!=nil{
		println(err)
		}
		firstname:=req.GetGreeting().GetFirstName()
		result +="hello"+firstname+"!!"


	}
}
func (*server)GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error{
	fmt.Println("GreetEveryone is invoked biDirectional")

	for{
		call,err:=stream.Recv()
		if err==io.EOF{
			return nil
		}
		if err != nil {
			fmt.Println("ERROR while birectional Stream in client request")

		}
		firstName:=call.GetGreeting().FirstName
		secondName:=call.GetGreeting().LastName
		result:= firstName+"ehyyy"+secondName
		sendErr:=stream.Send(&greetpb.GreetEveryoneResponse{
			Result: result,
		})
		if sendErr!=nil {
			fmt.Println("ERROR while birectional Stream in server Response")
		}





	}

}


func main()  {
	fmt.Println("heyy")
	lis,err:=net.Listen("tcp","0.0.0.0:5855")
	if err!=nil{
		log.Fatal("Failed,",err)
	}
	s:=grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s,&server{})
	if err:=s.Serve(lis); err!=nil{
		log.Fatal("failed ",err)
	}

}
