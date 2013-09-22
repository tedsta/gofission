package event

type Type uint8

var nextId Type

func NextId() Type {
	defer func() { nextId++ }()
	return nextId
}

type Event interface {
	Type() Type // Returns the event type for this event
}
