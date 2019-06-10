package main

import (
	"context"
	"log"
	"net"
	"strconv"

	pb "github.com/nicewook/grpc-go-quick-start/hi_pb"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) SayHi(ctx context.Context, in *pb.HiRequest) (*pb.HiReply, error) {
	log.Printf("Received: %v", in.Name)

	replyMsg := "Hi " + in.Name
	log.Printf("Send back to client: %s", replyMsg)
	return &pb.HiReply{Message: replyMsg}, nil
}

func (s *server) CountHi(ctx context.Context, in *pb.HiRequest) (*pb.HiReply, error) {
	log.Printf("Received: %v, length: %d", in.Name, len(in.Name))

	replyMsg := "Count " + in.Name + " length: " + strconv.Itoa(len(in.Name))
	log.Printf("Send back to client: %s", replyMsg)
	return &pb.HiReply{Message: replyMsg}, nil
}

func main() {
	log.Printf("grpc server start at port %s", port)

	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterHiServer(s, &server{})
	if err := s.Serve(l); err != nil {
		log.Fatalf("fail to serve: %v", err)
	}
}
