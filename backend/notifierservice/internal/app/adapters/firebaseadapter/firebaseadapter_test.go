package firebaseadapter_test

import (
	"fmt"
	"testing"

	"github.com/omerbeden/event-mate/backend/notifierservice/internal/app/adapters/firebaseadapter"
	"github.com/omerbeden/event-mate/backend/notifierservice/internal/app/domain/model"
	"github.com/stretchr/testify/assert"
)

type mockNotifier struct {
}

func (mockNotifier *mockNotifier) Send(message *model.PushMessage) (*model.PushMessageResponse, error) {
	fmt.Println("Test notifier")
	return &model.PushMessageResponse{true, "messageID"}, nil
}

func TestSend(t *testing.T) {
	adapter := firebaseadapter.FirebaseAdapter{Notifier: &mockNotifier{}}
	message := &model.PushMessage{}
	response, err := adapter.Notifier.Send(message)

	assert.NotNil(t, response)
	assert.NoError(t, err)

}
