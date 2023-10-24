package grpc

import (
	"context"
	"log"
	"net"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/infrastructure/grpc/pb"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/services/repositories"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	eventRepoService    repositories.EventRepoService
	userRepoService     repositories.UserRepoService
	userStatRepoService repositories.UserStatRepoService
	pb.UnimplementedProfileServiceServer
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserEventRequest) (*pb.GetUserEventResponse, error) {
	user, err := s.userRepoService.GetUserByID(uint(req.GetUserId()))
	if err != nil {
		return nil, err
	}
	return &pb.GetUserEventResponse{
		User: &pb.User{
			UserId:   int32(1),
			Name:     user.Name,
			LastName: user.LastName,
			About:    user.About,
			Photo:    user.Photo,
			//AttandedEvents: user.AttandedEvents, //TODO: MAPPING
			Adress: &pb.UserProfileAdress{City: user.Adress.City},
		},
	}, nil
}

func StartGRPCServer() {
	// TODO get configuration via di
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("error while listening tcp: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterProfileServiceServer(s, &server{eventRepoService: repositories.EventRepoService{}, userRepoService: repositories.UserRepoService{}, userStatRepoService: repositories.UserStatRepoService{}})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve %v", err)
	}
}
