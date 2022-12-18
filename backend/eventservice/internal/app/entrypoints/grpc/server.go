package grpc

import (
	"context"
	"log"
	"net"

	"github.com/go-redis/redis/v8"
	cacheadapter "github.com/omerbeden/event-mate/backend/eventservice/internal/app/adapters/cacheAdapter"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/adapters/database"
	adapters "github.com/omerbeden/event-mate/backend/eventservice/internal/app/adapters/repo"
	commandhandler "github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/command_handler"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/commands"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/model"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/ports/repo"
	"github.com/omerbeden/event-mate/backend/eventservice/internal/infra/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
	repo        repo.Repository
	redisOption *redis.Options
	pb.UnimplementedEventServiceServer
}

var redisOption = cacheadapter.RedisOption()

func (s *server) CreateEvent(ctx context.Context, req *pb.CreateEventRequest) (*pb.CreateEventResponse, error) {
	event := model.Event{
		Title:     req.GetEvent().GetTitle(),
		Category:  req.GetEvent().GetCategory(),
		CreatedBy: model.User{},
		Location:  model.Location{City: req.GetEvent().GetLocation().City},
	}

	client := redis.NewClient(redisOption)
	createCmd := &commands.CreateCommand{
		Repo:  s.repo,
		Event: event,
		Redis: cacheadapter.NewRedisAdapter(client),
	}
	defer client.Close()

	createCmdResult, err := commandhandler.HandleCommand[bool](createCmd)
	if err != nil {
		return &pb.CreateEventResponse{
			Status:  createCmdResult,
			Message: "event could not created",
		}, status.Error(codes.Unknown, "event could not created")
	}

	return &pb.CreateEventResponse{
		Status:  createCmdResult,
		Message: "created",
	}, nil

}

func (s *server) GetEvent(ctx context.Context, req *pb.GetEventRequest) (*pb.GetEventResponse, error) {
	getCommand := &commands.GetCommand{
		Repo:      s.repo,
		EventID:   req.GetEventId(),
		EventCity: req.GetEventCity(),
	}

	commandResult, err := commandhandler.HandleCommand[model.Event](getCommand)
	if err != nil {
		return nil, status.Error(codes.NotFound, "not found") //TODO refactor error matching
	}
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
	client := redis.NewClient(redisOption)
	defer client.Close()
	if !cmdResult.CacheHit {
		createCacheCommand := &commands.CreateCacheCommand{
			Redis: cacheadapter.NewRedisAdapter(client),
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

func StartGRPCServer(redisOpt *redis.Options) {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("error while listening tcp: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterEventServiceServer(s, &server{
		repo: &adapters.EventRepository{
			DB: database.InitPostgressConnection(),
		},
		redisOption: redisOpt,
	})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve %v", err)
	}
}
