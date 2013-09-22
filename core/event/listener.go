package event

type Listener interface {
	Listen(ch chan Event)
}
