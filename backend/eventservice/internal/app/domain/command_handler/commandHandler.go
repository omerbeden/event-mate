package commandhandler

import (
	"github.com/omerbeden/event-mate/backend/eventservice/internal/app/domain/commands"
)

func HandleCommand[T any](c commands.Command[T]) (T, error) {
	return c.Handle()
}
