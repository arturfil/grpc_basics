package main

import (
	"context"
	"log"

	pb "github.com/arturfil/first_grpc/sum/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
    log.Printf("Sum function was invoked %v\n", in)
    sum := in.A + in.B
    return &pb.SumResponse {
        Result: sum, 
    }, nil
}


