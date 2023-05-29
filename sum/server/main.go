package main

import (
	"log"
	"net"

	pb "github.com/arturfil/first_grpc/sum/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
    pb.SumServiceServer
}

func main() {
    lis, err := net.Listen("tcp", addr)
    if err != nil {
        log.Fatalf("Failed to listen on %v\n", err)
    }
    log.Printf("Listeningn on %s\n", addr)
    s := grpc.NewServer()
    pb.RegisterSumServiceServer(s, &Server{})

    if err = s.Serve(lis); err != nil {
        log.Fatal("Failed to serve %v\n", err)
    }
}
