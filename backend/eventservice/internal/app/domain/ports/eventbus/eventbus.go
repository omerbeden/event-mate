package eventbus

type EventBus interface {
	Subscribe()
	Publish()
}
