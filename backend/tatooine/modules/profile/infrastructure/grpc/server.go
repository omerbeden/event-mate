package grpc

import (
	"context"
	"log"
	"net"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/infrastructure/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedProfileServiceServer
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserEventRequest) (*pb.GetUserEventResponse, error) {
	return nil, nil
}

func StartGRPCServer() {
	// TODO get configuration via di
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("error while listening tcp: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterProfileServiceServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve %v", err)
	}
}
