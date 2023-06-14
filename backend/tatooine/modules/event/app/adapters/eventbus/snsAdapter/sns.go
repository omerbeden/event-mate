package snsadapter

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/ports/eventbus"
)

type SNSAdapter struct {
	Topic string
}

func PublishMessage(ctx context.Context, bus eventbus.SNSEventBus, input *sns.PublishInput) (*sns.PublishOutput, error) {
	return bus.Publish(ctx, input)
}

func SubscribeTopic(ctx context.Context, bus eventbus.SNSEventBus, input *sns.SubscribeInput) (*sns.SubscribeOutput, error) {
	return bus.Subscribe(ctx, input)
}

func (snsA *SNSAdapter) Publish(ctx context.Context, input *sns.PublishInput, optFns ...func(*sns.Options)) (*sns.PublishOutput, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := sns.NewFromConfig(cfg)

	result, err := PublishMessage(ctx, client, input)
	if err != nil {
		fmt.Println("err publish message")
		return nil, err
	}

	return result, err

}

func (snsA *SNSAdapter) Subscribe(ctx context.Context, input *sns.SubscribeInput, optFns ...func(*sns.Options)) (*sns.SubscribeOutput, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := sns.NewFromConfig(cfg)

	result, err := SubscribeTopic(context.Background(), client, input)
	if err != nil {
		fmt.Println("err subsribing to  topic")
		return nil, err
	}

	return result, err

}
