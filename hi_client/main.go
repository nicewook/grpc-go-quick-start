package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/nicewook/grpc-go-quick-start/hi_pb"
	"google.golang.org/grpc"
)

const (
	serverAddr  = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHiClient(conn)

	name := defaultName
	flag.Parse()
	if flag.NArg() > 0 {
		name = flag.Arg(0)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHi(ctx, &pb.HiRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting from the Server: %s", r.Message)

	r, err = c.CountHi(ctx, &pb.HiRequest{Name: name})
	if err != nil {
		log.Fatalf("could not count: %v", err)
	}
	log.Printf("Count from the Server: %s", r.Message)
}
