package chat

import (
	"golang.org/x/net/context"
	"learn-grpc/chat/chatpb"
	"log"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *chatpb.Message) (*chatpb.Message, error) {
	log.Printf("Receive message body from client: %s", message.Body)
	return &chatpb.Message{Body: "Hello From the Server!"}, nil
}
