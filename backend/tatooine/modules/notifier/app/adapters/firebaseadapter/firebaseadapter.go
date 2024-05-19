package firebaseadapter

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/notifier/app/domain/model"
)

type FirebaseNotifier struct {
	receiverId int64
}

func NewFirebaseNotifier(receiverId int64) *FirebaseNotifier {
	return &FirebaseNotifier{
		receiverId: receiverId,
	}
}

func (FirebaseNotifier *FirebaseNotifier) Send(message *model.PushMessage) (*model.PushMessageResponse, error) {
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: "gowith-1cdc5"}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		return &model.PushMessageResponse{Success: false, MessageId: "messageid"}, err
	}

	fcmClient, err := app.Messaging(context.Background())

	if err != nil {
		return &model.PushMessageResponse{Success: false, MessageId: "messageid"}, err
	}

	deviceToken, err := FirebaseNotifier.getDeviceToken()
	if err != nil {
		return &model.PushMessageResponse{Success: false, MessageId: "messageid"}, err
	}
	response, err := fcmClient.Send(context.Background(), &messaging.Message{

		Notification: &messaging.Notification{
			Title: message.Title,
			Body:  message.Body,
		},
		Token: deviceToken,
	})

	if err != nil {
		return &model.PushMessageResponse{Success: false, MessageId: "messageid"}, err
	}

	return &model.PushMessageResponse{Success: true, MessageId: response}, nil
}

func (FirebaseNotifier *FirebaseNotifier) getDeviceToken() (string, error) {
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: "gowith-1cdc5"}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	docResult, err := FirebaseNotifier.docAsMap(ctx, client)
	if err != nil {
		return "", err
	}

	return docResult["token"].(string), nil
}

func (FirebaseNotifier *FirebaseNotifier) docAsMap(ctx context.Context, client *firestore.Client) (map[string]interface{}, error) {
	diter := client.Collection("users").Doc(fmt.Sprint(FirebaseNotifier.receiverId)).Collection("device_tokens").Documents(ctx)
	docs, err := diter.GetAll()
	if err != nil {
		return nil, err
	}

	if len(docs) == 0 {
		return nil, fmt.Errorf("no documents found")
	}
	m := docs[0].Data()
	fmt.Printf("Document data: %#v\n", m)

	return m, nil
}
