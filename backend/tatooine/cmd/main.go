package main

import (
	"context"
	"log"
	"net"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omerbeden/event-mate/backend/tatooine/infra/grpc/pb"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/adapters/redisadapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/adapters/repo"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/entrypoints"
	postgres "github.com/omerbeden/event-mate/backend/tatooine/pkg/database"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
	EventService entrypoints.EventService
	pb.UnimplementedEventServiceServer
}

var redisOption = redisadapter.RedisOption()

func (s *server) CreateEvent(ctx context.Context, req *pb.CreateEventRequest) (*pb.CreateEventResponse, error) {

	event := model.Event{
		Title:     req.GetEvent().GetTitle(),
		Category:  req.GetEvent().GetCategory(),
		CreatedBy: model.User{ID: int64(req.GetUserId())},
		Location:  model.Location{City: req.GetEvent().GetLocation().City},
	}

	result, err := s.EventService.CreateEvent(ctx, event)

	if err != nil {
		return &pb.CreateEventResponse{
			Status:  result,
			Message: "event could not created",
		}, status.Errorf(codes.Unknown, "event could not created %s", err.Error())
	}
	return &pb.CreateEventResponse{
		Status:  result,
		Message: "created",
	}, nil

}

func (s *server) GetEvent(ctx context.Context, req *pb.GetEventRequest) (*pb.GetEventResponse, error) {

	result, err := s.EventService.GetEventById(ctx, req.GetEventId())

	return &pb.GetEventResponse{
		Event: &pb.Event{Id: result.ID,
			Title:    result.Title,
			Category: result.Category,
			Location: &pb.Location{City: result.Location.City}},
	}, err

}

func (s *server) GetFeed(ctx context.Context, req *pb.GetFeedByLocationRequest) (*pb.GetFeedByLocationResponse, error) {
	return nil, nil
}

func StartGRPCServer(redisOpt *redis.Options) {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("error while listening tcp: %v", err)
	}

	s := grpc.NewServer()
	dbPool := postgres.NewConn(&postgres.PostgresConfig{
		ConnectionString: "postgres://postgres:password@localhost:5432/test",
		Config:           pgxpool.Config{MinConns: 5, MaxConns: 10}})
	pb.RegisterEventServiceServer(s, &server{
		EventService: entrypoints.EventService{
			EventRepository:   repo.NewEventRepo(dbPool),
			LocationReposiroy: repo.NewLocationRepo(dbPool),
			RedisClient:       *redis.NewClient(redisOption),
		},
	})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve %v", err)
	}
}
