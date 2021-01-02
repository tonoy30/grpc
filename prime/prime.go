package prime

import (
	"learn-grpc/prime/primepb"
)

type Server struct {
}

func (s *Server) GetDividends(prime *primepb.Prime, stream primepb.PrimeServer_GetDividendsServer) error {
	divisor, n := 2, int(prime.GetNumber())
	for n > 1 {
		if n%divisor == 0 {
			stream.Send(&primepb.Response{
				Dividends: int32(divisor),
			})
			n /= divisor
		} else {
			divisor += 1
		}
	}
	return nil
}
