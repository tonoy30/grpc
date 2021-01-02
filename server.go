package main

import (
	"google.golang.org/grpc"
	"learn-grpc/chat"
	"learn-grpc/chat/chatpb"
	"learn-grpc/sum"
	"learn-grpc/sum/sumpb"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := chat.Server{}
	ss := sum.Server{}

	grpcServer := grpc.NewServer()

	chatpb.RegisterChatServiceServer(grpcServer, &s)
	sumpb.RegisterSumServiceServer(grpcServer, &ss)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}
