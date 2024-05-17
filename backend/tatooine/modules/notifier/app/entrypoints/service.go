package entrypoints

import (
	"github.com/omerbeden/event-mate/backend/tatooine/modules/notifier/app/adapters/firebaseadapter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/notifier/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/notifier/app/domain/port"
)

type NotifierService struct {
	notifier port.PushNotifier
}

func NewNotifierService() *NotifierService {
	return &NotifierService{
		notifier: &firebaseadapter.FirebaseNotifier{},
	}
}

func (NotifierService *NotifierService) Send(message *model.PushMessage) (*model.PushMessageResponse, error) {
	return NotifierService.notifier.Send(message)
}
