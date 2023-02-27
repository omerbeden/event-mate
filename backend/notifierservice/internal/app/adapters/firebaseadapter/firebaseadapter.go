package firebaseadapter

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/omerbeden/event-mate/backend/notifierservice/internal/app/domain/model"
	"github.com/omerbeden/event-mate/backend/notifierservice/internal/app/domain/port/pushnotification"
	"google.golang.org/api/option"
)

type FirebaseAdapter struct {
	Notifier pushnotification.PushNotifier
}

type FirebaseNotifier struct {
}

func (FirebaseNotifier *FirebaseNotifier) Send(message *model.PushMessage) (*model.PushMessageResponse, error) {
	opts := []option.ClientOption{option.WithCredentialsJSON([]byte{})}

	app, err := firebase.NewApp(context.Background(), nil, opts...)
	if err != nil {
		return &model.PushMessageResponse{Success: false, MessageId: "messageid"}, err
	}

	fcmClient, err := app.Messaging(context.Background())

	if err != nil {
		return &model.PushMessageResponse{Success: false, MessageId: "messageid"}, err
	}

	response, err := fcmClient.Send(context.Background(), &messaging.Message{

		Notification: &messaging.Notification{
			Title: message.Title,
			Body:  message.Body,
		},
		Token: message.DeviceToken, // it's a single device token
	})

	if err != nil {
		return &model.PushMessageResponse{Success: false, MessageId: "messageid"}, err
	}

	return &model.PushMessageResponse{Success: true, MessageId: response}, nil
}
