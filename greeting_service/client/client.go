package main

import (
	"GoGRPCTest/greeting_service/greetpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
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
	doServerStreaming(c)
	//doUnary(c)
}

func doUnary(c greetpb.DummyServiceClient){
	fmt.Println("Starting Unary call")
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

func doServerStreaming(c greetpb.DummyServiceClient){
	fmt.Println("Starting Server Streaming call")
	greet := &greetpb.Greeting{
		FirstName: "David",
		LastName:  "Ilenwabor",
		Age:       32,
	}
	request := greetpb.GreetRequest{Greeting: greet}
	resStream, err := c.GreetManyTimesFromServer(context.Background(), &request)
	if err != nil{
		log.Fatalf("Could not call server to stream greeting, an error occured %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF{ //this error is returned when a stream closes
			log.Printf("Streaming ended")
			break
		}
		if err != nil{
			log.Fatalf("Error while receiving stream %v", err)
		}
		log.Printf(msg.GetGreeting())
	}
}
