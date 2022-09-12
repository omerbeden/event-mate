package grpc

import (
	"context"
	"log"
	"net"

	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/ports/repo"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/infra/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	repo repo.Repository
	pb.UnimplementedEventServiceServer
}

func (s *server) CreateEvent(ctx context.Context, req *pb.CreateEventRequest) (*pb.CreateEventResponse, error) {
	//TODO: istead of mapping here , maybe commands can be used here to map grpc event to domain event and call repo.CreateEvent correctly
	//s.repo.CreateEvent(req.Event)
	return nil, nil
}
func (s *server) GetEvent(ctx context.Context, req *pb.GetEventRequest) (*pb.GetEventResponse, error) {

	return nil, nil
}

func StartGRPCServer() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("error while listening tcp: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterEventServiceServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve %v", err)
	}
}
