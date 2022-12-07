package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "grpc-practice/proto/echo"

	_ "google.golang.org/grpc/encoding/gzip"
)

var port = flag.Int("port", 50002, "the port to serve on")

func main() {
	flag.Parse()
	fmt.Printf("server starting on port %d...\n", *port)

	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &ecServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type ecServer struct {
	pb.UnimplementedEchoServer
}

func (s *ecServer) UnaryEcho(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{Message: req.Message}, nil
}
