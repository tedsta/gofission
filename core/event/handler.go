package event

type Handler interface {
	HandleEvent(e Event)
}
