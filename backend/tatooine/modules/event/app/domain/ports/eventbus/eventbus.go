package eventbus

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type SNSEventBus interface {
	Subscribe(ctx context.Context, params *sns.SubscribeInput, optFns ...func(*sns.Options)) (*sns.SubscribeOutput, error)
	Publish(ctx context.Context, input *sns.PublishInput, optFns ...func(*sns.Options)) (*sns.PublishOutput, error)
}
