package greet

import (
	"context"
	"fmt"
	"learn-grpc/greet/greetpb"
	"log"
	"time"
)

type Server struct {
}

func (s *Server) Greeting(ctx context.Context, greet *greetpb.Greet) (*greetpb.Greet, error) {
	log.Printf("Receive greet message from client: %s", greet.Message)
	return &greetpb.Greet{Message: fmt.Sprintf("Hello %s", greet.Message)}, nil
}
func (s *Server) GreetManyTimes(greet *greetpb.Greet, stream greetpb.GreetService_GreetManyTimesServer) error {
	for i := 0; i < 10; i++ {
		res := &greetpb.Greet{Message: fmt.Sprintf("%s %d\n", "Message", i)}
		stream.SendMsg(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}
