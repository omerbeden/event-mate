package pushnotification

import "github.com/omerbeden/event-mate/backend/notifierservice/internal/app/domain/model"

type PushNotifier interface {
	Send(message *model.PushMessage) (*model.PushMessageResponse, error)
}
