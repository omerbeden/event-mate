package firebaseadapter

import (
	"context"
	"encoding/json"
	"fmt"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/notifier/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/notifier/app/domain/port/db"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/notifier/app/domain/port/pushnotification"
	"google.golang.org/api/option"
)

type FirebaseAdapter struct {
	Notifier pushnotification.PushNotifier
	Db       db.Database
}

type FirebaseNotifier struct{}

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

type FirebaseDB struct{}

func (FirebaseDB *FirebaseDB) PersistDeviceToken(token model.Token) (bool, error) {

	ctx := context.Background()
	opts := []option.ClientOption{option.WithCredentialsJSON([]byte{})}
	app, err := firebase.NewApp(ctx, nil, opts...)

	if err != nil {
		return false, fmt.Errorf("app cannot created")
	}

	dbClient, err := app.Database(ctx)
	if err != nil {
		return false, fmt.Errorf("db client cannot created")
	}

	ref := dbClient.NewRef("test/test")

	json, err := json.Marshal(token)
	if err != nil {
		return false, fmt.Errorf("marshal err")
	}

	err1 := ref.Set(ctx, json)

	if err1 != nil {
		return false, fmt.Errorf("unnkow err")
	}

	return true, nil

}

func (FirebaseDB *FirebaseDB) GetToken(userId string) (*model.Token, error) {

	ctx := context.Background()
	opts := []option.ClientOption{option.WithCredentialsJSON([]byte{})}
	app, err := firebase.NewApp(ctx, nil, opts...)

	if err != nil {
		return nil, fmt.Errorf("app cannot created")
	}

	dbClient, err := app.Database(ctx)
	if err != nil {
		return nil, fmt.Errorf("db client cannot created")
	}

	ref := dbClient.NewRef("test/test")

	var token model.Token
	if err := ref.Get(ctx, &token); err != nil {
		return nil, fmt.Errorf("error in reading token")
	}

	return &token, nil

}

//TODO: Birde update lazım ,  frontend de refleshment için
