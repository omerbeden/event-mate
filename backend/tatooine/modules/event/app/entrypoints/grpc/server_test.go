package grpc

import (
	"context"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/infra/grpc/pb"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

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
			Title:    "End2End Test",
			Category: "Rock",
			Location: &pb.Location{
				City: "Sakarya",
			},
		},
	})

	t.Log("endpoint called")
	expected := &pb.CreateEventResponse{
		Message: "created",
		Status:  true,
	}

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.EqualValues(t, expected.Message, response.Message)
	assert.EqualValues(t, expected.Status, response.Status)

}

func TestGetEvent(t *testing.T) {
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

	getRequest := pb.GetEventRequest{
		EventId:   "1",
		EventCity: "Sakarya",
	}

	t.Log("endpoint calling")
	response, err := client.GetEvent(context.Background(), &getRequest)
	if err != nil {
		panic(err)
	}

	t.Log("endpoint called")

	expected := &pb.GetEventResponse{
		Event: &pb.Event{
			Id:       "1",
			Title:    "End2End Test",
			Category: "Rock",
			Location: &pb.Location{
				City: "Sakarya",
			},
		},
	}

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, expected.GetEvent().Category, response.GetEvent().Category)
	assert.Equal(t, expected.GetEvent().GetLocation().City, response.GetEvent().GetLocation().City)
	assert.Equal(t, expected.GetEvent().Title, response.GetEvent().Title)
}
