package main

import (
	"log"
	"net"

	context "golang.org/x/net/context"

	"github.com/chapin/microservices-k8s-grpc-go/pb"

	"google.golang.org/grpc"

	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGCGServiceServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) Compute(ctx context.Context, r *pb.GCDRequest) (*pb.GCDResponse, error) {
	a, b := r.A, r.B
	if b != 0 {
		a, b = b, a%b
	}
	return &pb.GCDResponse{Result: a}, nil
}
