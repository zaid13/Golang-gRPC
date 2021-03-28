package main


import (
	"context"
	"fmt"
	"gRPC/calculator/calculatorpb"


	//"gRPC/greet/greetpb"

	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
}
func (*server)Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error){
	fmt.Println("calculator function was invoked")


	firstnumber := req.FirstNumber
	secondnumber:= req.SecondNumber
	fmt.Println("FirstNumber was + "+string(firstnumber))
	fmt.Println("SecondNumber was + "+string(secondnumber))

	result:= firstnumber+secondnumber
	res:=&calculatorpb.SumResponse{
		SumResult: result,
	}
	return res,nil

}

func (*server)PrimeNumberDecomposition( req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error{
	fmt.Println("PrimeNumberDecomposition function was invoked",req)
	number:= req.GetNumber()
	divisor:=int64(2)

	for number >1 {
		if number%divisor ==0{
			stream.Send(&calculatorpb.PrimeNumberDecompositionResponse{
				Primefactor: divisor,
			})
			number = number/divisor

		} else {
			divisor++
			fmt.Println("divisor has increased to ",divisor)
					}
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
	calculatorpb.RegisterCalculatorServiceServer(s,&server{})
	if err:=s.Serve(lis); err!=nil{
		log.Fatal("failed ",err)
	}

}
