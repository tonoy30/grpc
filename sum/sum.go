package sum

import (
	"golang.org/x/net/context"
	"learn-grpc/sum/sumpb"
	"log"
)

type Server struct {
}

func (c *Server) Sum(ctx context.Context, input *sumpb.Input) (*sumpb.Result, error) {
	log.Printf("Recevied input from client: %d, %d\n", input.Num1, input.Num2)
	return &sumpb.Result{Ans: input.Num1 + input.Num2}, nil
}
