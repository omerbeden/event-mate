package grpc

import (
	"context"
	"log"
	"net"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
	cacheadapter "github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/adapters/cacheAdapter"
	adapters "github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/adapters/repo"
	commandhandler "github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/command_handler"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/commands"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/model"
	repo "github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/ports/repo"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/infra/grpc/pb"
	postgres "github.com/omerbeden/event-mate/backend/tatooine/pkg/database"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
	eventRepo    repo.EventRepository
	locationRepo repo.LocationRepository
	redisOption  *redis.Options
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
		Repo:  s.eventRepo,
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
		Repo:      s.eventRepo,
		EventID:   req.GetEventId(),
		EventCity: req.GetEventCity(),
	}

	commandResult, err := commandhandler.HandleCommand[*model.Event](getCommand)
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
		Repo:     s.eventRepo,
		Location: location,
	}

	cmdResult, err := commandhandler.HandleCommand[*model.GetFeedCommandResult](getFeedCommand)
	client := redis.NewClient(redisOption)
	defer client.Close()
	if !cmdResult.CacheHit {
		createCacheCommand := &commands.CreateCacheCommand{
			Redis: cacheadapter.NewRedisAdapter(client),
			Key:   location.City,
			Posts: cmdResult.Events,
		}
		_, createErr := commandhandler.HandleCommand[bool](createCacheCommand)
		if createErr != nil {
			return nil, err
		}
	}

	var events []*pb.Event
	for i := range cmdResult.Events {
		events[i] = &pb.Event{
			Id:       strconv.FormatUint(uint64(cmdResult.Events[i].ID), 10),
			Title:    cmdResult.Events[i].Title,
			Category: cmdResult.Events[i].Category,
			Location: &pb.Location{City: cmdResult.Events[i].Location.City},
		}
	}

	return &pb.GetFeedByLocationResponse{
		Event: events,
	}, err

}

func StartGRPCServer(redisOpt *redis.Options) {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("error while listening tcp: %v", err)
	}

	s := grpc.NewServer()
	dbPool := postgres.NewConn(&postgres.PostgresConfig{ConnectionString: "", Config: pgxpool.Config{}})
	pb.RegisterEventServiceServer(s, &server{
		eventRepo:    adapters.NewEventRepo(dbPool),
		locationRepo: adapters.NewLocationRepo(dbPool),
		redisOption:  redisOpt,
	})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve %v", err)
	}
}
