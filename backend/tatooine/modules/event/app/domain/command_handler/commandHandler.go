package commandhandler

import (
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/commands"
)

func HandleCommand[T any](c commands.Command[T]) (T, error) {
	return c.Handle()
}
