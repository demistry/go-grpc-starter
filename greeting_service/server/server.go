package main

import (
	"GoGRPCTest/greeting_service/greetpb"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	PORT       = "50051"
	TcpAddress = "0.0.0.0:"
)

type server struct {

}

func (s server) Greet(ctx context.Context, request *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {

	return nil, nil
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
