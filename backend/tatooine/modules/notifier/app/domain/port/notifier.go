package port

import "github.com/omerbeden/event-mate/backend/tatooine/modules/notifier/app/domain/model"

type PushNotifier interface {
	Send(message *model.PushMessage) (*model.PushMessageResponse, error)
}
