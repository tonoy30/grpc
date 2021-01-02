package greet

import (
	"fmt"
	"golang.org/x/net/context"
	"learn-grpc/greet/greetpb"
	"log"
)

type Server struct {
}

func (s *Server) Greeting(ctx context.Context, greet *greetpb.Greet) (*greetpb.Greet, error) {
	log.Printf("Receive greet message from client: %s", greet.Message)
	return &greetpb.Greet{Message: fmt.Sprintf("Hello %s", greet.Message)}, nil
}
