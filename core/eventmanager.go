package core

var NextEventId int

type Event interface {
	Type() int // Returns the event type for this event
}

type EventHandler interface {
	HandleEvent(event Event)
}

type EventManager struct {
	receivers [][]EventHandler
}

// NewEventManager creates a new event manager
func NewEventManager(eventCount int) *EventManager {
	return &EventManager{make([][]EventHandler, eventCount)}
}

// AddHandler adds an event handler for a particular event type
func (e *EventManager) AddHandler(eventType int, handler EventHandler) {
	// Make sure it's a valid event type
	if eventType < int(len(e.receivers)) {
		e.receivers[eventType] = append(e.receivers[eventType], handler)
	}
}

// FireEvent fires an event to all receivers receiving
func (e *EventManager) FireEvent(event Event) {
	for _, receiver := range e.receivers[event.Type()] {
		go receiver.HandleEvent(event) // Why wait for events to get handled?
	}
}