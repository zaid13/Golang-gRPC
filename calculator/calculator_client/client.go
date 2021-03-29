package main

import (
	"context"
	"fmt"
	"gRPC/calculator/calculatorpb"
	"io"

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
	//doClientStreaming(c)
	doBiDirectional(c)

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

func doClientStreaming(c calculatorpb.CalculatorServiceClient)  {
	fmt.Println("started Client streaming  gRPC ")

	requests:=[]float64{
				3,4,5,
	}

	stream, err:=c.ComputeAverage(context.Background())
	if err!=nil{
		fmt.Println("error senfning streaming request to server ")
	}
	for _,val:= range requests {
		fmt.Println("sending ",val)

		stream.Send(&calculatorpb.ComputeAverageRequest{
			Number: val,
		})
	}
	res,err :=stream.CloseAndRecv()
	if err != nil {
		println("error receiving message form server",err)
	}

	fmt.Println("SUCCESSFull received message ",res.GetAverage())


}
func doBiDirectional(c calculatorpb.CalculatorServiceClient)  {
	println("doing bidirectrional")

	stream , errr:=c.FindMaximum(context.Background())
	if errr != nil {
		log.Fatalf("eRRRO WHILE OPENIGN STREAM")
	}
	waitc :=make(chan struct {})

	go func() {
		numbers:=[]float64{4,7,2,19,6,32}
		for _,num:=range numbers{
			fmt.Println("senfing numbers to server")
			stream.Send(&calculatorpb.FindMaximumRequest{
				Number: num,
			})

		}
		stream.CloseSend()

	}()
	go func() {

		for {
			fmt.Println("reciving max from  server")
			res, err:=stream.Recv()
			if err ==io.EOF {
				break

			}
			if err != nil {
				log.Fatalf("errro fetching response from server ",err)
				break
			}
			fmt.Println("recivbed number from server ",res.GetMax())
		}

		close(waitc)
	}()

<-waitc
}