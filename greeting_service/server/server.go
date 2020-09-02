package main

import (
	"GoGRPCTest/greeting_service/greetpb"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const (
	PORT       = "50051"
	TcpAddress = "0.0.0.0:"
)

type server struct {

}

func (s server) Greet(ctx context.Context, request *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	firstName := request.GetGreeting().FirstName
	lastName := request.GetGreeting().LastName
	age := request.GetGreeting().Age
	greeting := fmt.Sprintf("Hello %s %s, you are %d years old",firstName,lastName, age)
	response := &greetpb.GreetResponse{
		Status: true,
		Message:  "Successfully returned greet response",
		Greeting: greeting,
	}
	return response, nil
}

func (s server) GreetManyTimesFromServer(request *greetpb.GreetRequest, stream greetpb.DummyService_GreetManyTimesFromServerServer) error {
	firstName := request.GetGreeting().FirstName
	lastName := request.GetGreeting().LastName
	age := request.GetGreeting().Age

	for i := 0; i < 20; i++{
		greeting := fmt.Sprintf("Hello %s %s, you are %d years old. Enjoy greeting %d", firstName, lastName, age, i)
		response := &greetpb.GreetManyTimesResponse{
			Status:   true,
			Message:  "New greeting available",
			Greeting: greeting,
		}
		_ = stream.Send(response)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func main(){

	lis, err := net.Listen("tcp", TcpAddress + PORT)
	if err != nil{
		log.Fatal("Could not initialize listener due to ", err.Error())
	}
	fmt.Println("Running server on port: 50051")
	s := grpc.NewServer()
	greetpb.RegisterDummyServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server %v", err)
	}
}
