package main


import (
	"context"
	"fmt"
	"gRPC/greet/greetpb"
	"google.golang.org/grpc"
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
