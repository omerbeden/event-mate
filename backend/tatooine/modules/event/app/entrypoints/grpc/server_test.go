package grpc

import (
	"context"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/infra/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ExceptedResponse struct {
	message string
	status  bool
}

func TestCreateEvent(t *testing.T) {
	go func() {
		redisOpt := &redis.Options{
			Addr:     "Localhost:6379",
			Password: "",
			DB:       0,
		}
		StartGRPCServer(redisOpt)
	}()

	t.Log("server created\n")

	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewEventServiceClient(conn)
	response, err := client.CreateEvent(context.Background(), &pb.CreateEventRequest{
		UserId: "1",
		Event: &pb.Event{
			Id:       "1",
			Title:    "End2End Test",
			Category: "Rock",
			Location: &pb.Location{
				City: "Sakarya",
			},
		},
	})

	t.Log("endpoint called")

	if err != nil {
		panic(err)
	}

	expected := &ExceptedResponse{
		message: "created",
		status:  true,
	}

	if expected.message != response.GetMessage() && expected.status != response.GetStatus() {
		t.Errorf("got %q want %q", expected.message, response.Message)
		t.Errorf("got %t want %t", expected.status, response.Status)

	}

	t.Log(response)
}
