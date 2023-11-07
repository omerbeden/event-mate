package command

type Command[T any] interface {
	Handle() (T, error)
}

func HandleCommand[T any](c Command[T]) (T, error) {
	return c.Handle()
}
