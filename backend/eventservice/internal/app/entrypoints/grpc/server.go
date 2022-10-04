package grpc

import (
	"context"
	"log"
	"net"

	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/adapters/database"
	adapters "github.com/omerbeden/event-mate/backend/eventservice/internal/app/adapters/repo"
	commandhandler "github.com/omerbeden/event-mate/backend/eventservice/internal/app/command_handler"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/commands"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/model"
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
	createCommand := &commands.CreateCommand{
		Repo:  s.repo,
		Event: req.GetEvent(),
	}
	commandResult, err := commandhandler.HandleCommand[bool](createCommand)

	return &pb.CreateEventResponse{
		Status:  commandResult,
		Message: "created",
	}, err

}

func (s *server) GetEvent(ctx context.Context, req *pb.GetEventRequest) (*pb.GetEventResponse, error) {
	getCommand := &commands.GetCommand{
		Repo:    s.repo,
		EventID: req.GetEventId(),
	}

	commandResult, err := commandhandler.HandleCommand[model.Event](getCommand)
	return &pb.GetEventResponse{
		Event: &pb.Event{Title: commandResult.Title, Category: commandResult.Category},
	}, err

}

func StartGRPCServer() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("error while listening tcp: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterEventServiceServer(s, &server{
		repo: &adapters.EventRepository{
			DB: database.InitPostgressConnection(),
		},
	})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve %v", err)
	}
}
