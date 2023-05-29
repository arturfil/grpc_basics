package main

import (
	"context"
	"log"

	pb "github.com/arturfil/first_grpc/greet/proto"
)

func  doGreet(c pb.GreetServiceClient) {
    log.Println("doGreet was invoked")
    res, err := c.Greet(context.Background(), &pb.GreetRequest {
        FirstName: "Arturo",
    })

    if err != nil {
        log.Fatalf("Could no greet: %v\n", err)
    }

    log.Printf("Greeting: %s\n", res.Result)
}
