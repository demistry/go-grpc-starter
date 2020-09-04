package main

import (
	"GoGRPCTest/greeting_service/greetpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
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
	//doUnary(c)
	//doServerStreaming(c)
	doClientStreaming(c)

}

//Unary RPC call
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

//Server streaming RPC call
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

func doClientStreaming(c greetpb.DummyServiceClient){
	fmt.Println("Starting client streaming call")

	requests := []*greetpb.GreetRequest{
		&greetpb.GreetRequest{Greeting: &greetpb.Greeting{
			FirstName: "David",
			LastName:  "Ilenwabor",
			Age:       50,
		}},
		&greetpb.GreetRequest{Greeting: &greetpb.Greeting{
			FirstName: "David",
			LastName:  "Oshioke",
			Age:       100,
		}},
		&greetpb.GreetRequest{Greeting: &greetpb.Greeting{
			FirstName: "Davidemi",
			LastName:  "Ilenz",
			Age:       25,
		}},
	}
	stream,err := c.LongGreetFromClient(context.Background()) //since its a stream, you dont need to pass in the request here
	if err != nil{
		log.Fatalf("Error occured with client streaming %v", err)
	}
	for _,req := range requests {
		fmt.Printf("Sending request %v\n", req)
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil{
		log.Fatalf("Error while receiving response after streaming %v", err)
	}
	log.Printf("Response received for client streaming is %v", res)
}
