package command

type Command[T any] interface {
	Handle() (T, error)
}
