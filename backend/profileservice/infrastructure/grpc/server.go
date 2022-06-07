package grpc

import (
	"log"
	"net"

	"github.com/omerbeden/event-mate/backend/profileservice/services/repositories"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	eventRepoService    repositories.EventRepoService
	userRepoService     repositories.UserRepoService
	userStatRepoService repositories.UserStatRepoService
}

func StartGRPCServer() {
	// TODO get configuration via di
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("error while listening tcp: %v", err)
	}

	s := grpc.NewServer()
	//TODO create pb
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve %v", err)
	}
}
