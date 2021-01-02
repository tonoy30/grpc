package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"learn-grpc/chat/chatpb"
	"learn-grpc/greet/greetpb"
	"learn-grpc/sum/sumpb"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := chatpb.NewChatServiceClient(conn)
	s := sumpb.NewSumServiceClient(conn)
	gs := greetpb.NewGreetServiceClient(conn)

	response, err := c.SayHello(context.Background(), &chatpb.Message{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s\n", response.Body)

	ans, err := s.Sum(context.Background(), &sumpb.Input{Num1: 2, Num2: 3})
	if err != nil {
		log.Fatalf("Error when calling Sum: %s", err)
	}
	log.Printf("Response from sum server: %d\n", ans.Ans)

	greet, err := gs.Greeting(context.Background(), &greetpb.Greet{
		Message: "Hello ",
	})
	if err != nil {
		log.Fatalf("Error when calling Sum: %s", err)
	}
	log.Printf("Response from sum server: %s\n", greet.Message)
}
