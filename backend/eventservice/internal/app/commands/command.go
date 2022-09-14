package commands

// TODO: Generic olabilr, get event i desteklemesi için , create ve update lerdede true yada false dönebilir
type Command[T any] interface {
	Handle() (T, error)
}
