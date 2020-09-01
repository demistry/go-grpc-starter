package main

import (
	"GoGRPCTest/greeting_service/greetpb"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {

}

func main(){

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil{
		log.Fatal("Could not initialize listener due to ", err.Error())
	}
	fmt.Println("Running server on port: 50051")
	s := grpc.NewServer()
	greetpb.RegisterDummyServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to server %v", err)
	}
}
