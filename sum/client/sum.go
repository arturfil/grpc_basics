package main

import (
	"context"
	"log"

	pb "github.com/arturfil/first_grpc/sum/proto"
)

func doSum(c pb.SumServiceClient) {
    log.Println("doSum was invoked")
    a, b := 5, 8
    res, err := c.Sum(context.Background(), &pb.SumRequest{
        A: int32(a),
        B: int32(b),
    })
    log.Printf("Values are %d, %d", a, b)
    if err != nil {
        log.Fatalf("Could not sum values: %v\n", err)
    }
    log.Printf("Result: %d\n", res.Result)
}
