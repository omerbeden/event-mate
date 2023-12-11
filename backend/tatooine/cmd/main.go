package main

import (
	"context"
	"log"
	"net"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omerbeden/event-mate/backend/tatooine/infra/grpc/pb"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/adapters/redisadapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/adapters/repo"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/entrypoints"
	postgres "github.com/omerbeden/event-mate/backend/tatooine/pkg/database"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
	EventService entrypoints.ActivityService
	pb.UnimplementedActivityServiceServer
}

var redisOption = redisadapter.RedisOption()

func (s *server) CreateEvent(ctx context.Context, req *pb.CreateActivityRequest) (*pb.CreateActivityResponse, error) {

	event := model.Event{
		Title:     req.GetActivity().GetTitle(),
		Category:  req.GetActivity().GetCategory(),
		CreatedBy: model.User{ID: int64(req.GetUserId())},
		Location:  model.Location{City: req.GetActivity().GetLocation().City},
	}

	result, err := s.EventService.CreateActivity(ctx, event)

	if err != nil {
		return &pb.CreateActivityResponse{
			Status:  result,
			Message: "event could not created",
		}, status.Errorf(codes.Unknown, "event could not created %s", err.Error())
	}
	return &pb.CreateActivityResponse{
		Status:  result,
		Message: "created",
	}, nil

}

func (s *server) GetEvent(ctx context.Context, req *pb.GetActivityByIdRequest) (*pb.GetActivityByIdResponse, error) {

	result, err := s.EventService.GetActivityById(ctx, req.GetActivityId())

	return &pb.GetActivityByIdResponse{
		Activity: &pb.Activity{Id: result.ID,
			Title:    result.Title,
			Category: result.Category,
			Location: &pb.Location{City: result.Location.City}},
	}, err

}

func (s *server) GetFeed(ctx context.Context, req *pb.GetActivitiesByLocationRequest) (*pb.GetActivitiesByLocationResponse, error) {
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
	pb.RegisterActivityServiceServer(s, &server{
		EventService: entrypoints.ActivityService{
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
