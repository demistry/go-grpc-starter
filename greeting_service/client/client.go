package main

import (
	"GoGRPCTest/greeting_service/greetpb"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main(){


	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) //to disable SSL, disable in production apps
	if err != nil{
		log.Fatal("Could not connect to server %v", err)
	}
	defer conn.Close()
	c := greetpb.NewDummyServiceClient(conn)
	fmt.Println("Created client %f ", c)
}
