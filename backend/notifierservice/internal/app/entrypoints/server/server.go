package server

import (
	"context"
	"log"
	"net"

	"github.com/omerbeden/event-mate/backend/notifierservice/internal/infra/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//this may have one andpoint that  read device token and store it

type server struct {
	pb.UnimplementedNotifierServiceServer
}

func (s *server) PersistDeviceToken(ctx context.Context, req *pb.PersistDeviceTokenRequest) (*pb.PersistDeviceTokenResponse, error) {

	//firebase connection
	return nil, nil
}

func StartGRPCServer() {
	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalf("errror while listening tcp: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterNotifierServiceServer(s, &server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve %v", err)
	}
}
