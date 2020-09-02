package main

import (
	"GoGRPCTest/greeting_service/greetpb"
	"context"
	"google.golang.org/grpc"
	"log"
)

const PORT = "50051"
func main(){


	//initializing withInsecure to disable SSL certification since it isnt needed, disable in production apps
	conn, err := grpc.Dial("localhost:" + PORT, grpc.WithInsecure())
	if err != nil{
		log.Fatalf("Could not connect to server %v", err)
	}
	defer conn.Close()
	c := greetpb.NewDummyServiceClient(conn)
	greet := &greetpb.Greeting{
		FirstName: "David",
		LastName:  "Ilenwabor",
		Age:       32,
	}
	request := greetpb.GreetRequest{Greeting: greet}
	resp, err := c.Greet(context.Background(), &request)
	if err != nil{
		log.Fatalf("Could not receive greeting, an error occured %v", err)
	}
	log.Printf("Response from greeting: %v\n", resp)
}
