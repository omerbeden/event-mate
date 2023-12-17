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

	event := model.Activity{
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

func (s *server) GetEventByID(ctx context.Context, req *pb.GetActivityByIdRequest) (*pb.GetActivityByIdResponse, error) {

	result, err := s.EventService.GetActivityById(ctx, req.GetActivityId())

	return &pb.GetActivityByIdResponse{
		Activity: &pb.Activity{Id: result.ID,
			Title:    result.Title,
			Category: result.Category,
			Location: &pb.Location{City: result.Location.City}},
	}, err

}

func (s *server) GetActivitiesByLocation(ctx context.Context, req *pb.GetActivitiesByLocationRequest) (*pb.GetActivitiesByLocationResponse, error) {

	location := model.Location{
		City: req.GetLocation().GetCity(),
	}

	activities, err := s.EventService.GetActivitiesByLocation(ctx, location)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "could not get activities by city %s", err.Error())
	}

	var activitiesPB []*pb.Activity
	for _, activity := range activities {
		activityPB := pb.Activity{
			Id:       activity.ID,
			Title:    activity.Title,
			Category: activity.Category,
			Location: &pb.Location{
				City: activity.Location.City,
			},
		}
		activitiesPB = append(activitiesPB, &activityPB)
	}

	return &pb.GetActivitiesByLocationResponse{
		Activity: activitiesPB,
	}, nil
}

func (server *server) AddParticipant(ctx context.Context, req *pb.AddParticipantRequest) (*pb.AddParticipantResponse, error) {

	participant := model.User{
		ID: req.GetParticipant().GetId(),
	}
	err := server.EventService.AddParticipant(participant, req.GetActivityId())
	if err != nil {
		return &pb.AddParticipantResponse{
			Status:  false,
			Message: "participant could not be added",
		}, status.Errorf(codes.Internal, "participant could not be added %s", err.Error())
	}

	return &pb.AddParticipantResponse{
		Status:  true,
		Message: "OK",
	}, nil
}

func (server *server) GetParticipants(ctx context.Context, req *pb.GetParticipantRequest) (*pb.GetParticipantResponse, error) {
	result, err := server.EventService.GetParticipants(req.GetActivityId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "participant could not be added %s", err.Error())
	}

	var userList []*pb.User
	for _, participants := range result {
		userList = append(userList, &pb.User{
			Id: participants.ID,
		})
	}
	return &pb.GetParticipantResponse{
		Users: userList,
	}, nil
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
			ActivityRepository: repo.NewActivityRepo(dbPool),
			LocationReposiroy:  repo.NewLocationRepo(dbPool),
			RedisClient:        *redis.NewClient(redisOption),
		},
	})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve %v", err)
	}
}
