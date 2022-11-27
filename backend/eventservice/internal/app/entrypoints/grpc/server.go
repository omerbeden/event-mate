package grpc

import (
	"context"
	"log"
	"net"

	cacheadapter "github.com/omerbeden/event-mate/backend/eventservice/internal/app/adapters/cacheAdapter"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/adapters/database"
	adapters "github.com/omerbeden/event-mate/backend/eventservice/internal/app/adapters/repo"
	commandhandler "github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/command_handler"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/commands"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/model"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/ports/repo"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/infra/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	repo repo.Repository
	pb.UnimplementedEventServiceServer
}

var redisOption = cacheadapter.RedisOption()

func (s *server) CreateEvent(ctx context.Context, req *pb.CreateEventRequest) (*pb.CreateEventResponse, error) {
	event := model.Event{
		Title:     req.GetEvent().GetTitle(),
		Category:  req.GetEvent().GetCategory(),
		CreatedBy: model.User{},
		Location:  model.Location{City: "sakarya"},
	}

	createCommand := &commands.CreateCommand{
		Repo:  s.repo,
		Event: event,
		Redis: cacheadapter.NewRedisAdapter(redisOption),
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

func (s *server) GetFeed(ctx context.Context, req *pb.GetFeedByLocationRequest) (*pb.GetFeedByLocationResponse, error) {
	location := &model.Location{
		City: req.GetLocation().GetCity(),
	}

	getFeedCommand := &commands.GetFeedCommand{
		Repo:     s.repo,
		Location: location,
	}

	cmdResult, err := commandhandler.HandleCommand[*model.GetFeedCommandResult](getFeedCommand)

	if !cmdResult.CacheHit {
		createCacheCommand := &commands.CreateCacheCommand{
			Redis: cacheadapter.NewRedisAdapter(redisOption),
			Key:   location.City,
			Posts: *cmdResult.Events,
		}
		_, createErr := commandhandler.HandleCommand[bool](createCacheCommand)
		if createErr != nil {
			return nil, err
		}
	}

	return nil, err

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
