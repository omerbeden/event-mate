package sns_test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go/aws"
	snsadapter "github.com/omerbeden/event-mate/backend/eventservice/internal/app/adapters/eventbus/snsAdapter"
	"github.com/stretchr/testify/assert"
)

type SNSMock struct{}

func (m *SNSMock) Publish(ctx context.Context, input *sns.PublishInput, optFns ...func(*sns.Options)) (*sns.PublishOutput, error) {
	output := &sns.PublishOutput{
		MessageId: aws.String("123"),
	}

	return output, nil
}

func (m *SNSMock) Subscribe(ctx context.Context, params *sns.SubscribeInput, optFns ...func(*sns.Options)) (*sns.SubscribeOutput, error) {
	output := &sns.SubscribeOutput{
		SubscriptionArn: aws.String("arn"),
	}
	return output, nil
}

func TestPublish(t *testing.T) {

	mock := &SNSMock{}
	input := &sns.PublishInput{
		Message:  aws.String("message"),
		TopicArn: aws.String("topic"),
	}

	resp, err := snsadapter.PublishMessage(context.Background(), mock, input)
	if err != nil {
		t.Log("Got an error ...:")
		t.Log(err)
		return
	}

	t.Log("Message ID: " + *resp.MessageId)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

}

func TestSubscribe(t *testing.T) {
	mock := &SNSMock{}
	input := &sns.SubscribeInput{
		Endpoint:              aws.String("test@mail"),
		Protocol:              aws.String("email"),
		ReturnSubscriptionArn: true,
		TopicArn:              aws.String("topic"),
	}

	resp, err := snsadapter.SubscribeTopic(context.Background(), mock, input)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
