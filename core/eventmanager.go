package core

import (
	"sync"
)

type EventType uint8

var nextEventId EventType

func NextEventId() EventType {
	defer func() { nextEventId++ }()
	return nextEventId
}

type Event interface {
	Type() EventType // Returns the event type for this event
}

type EventHandler interface {
	HandleEvent(event Event)
}

type EventManager struct {
	receivers [][]EventHandler
}

// AddHandler adds an event handler for a particular event type
func (e *EventManager) AddHandler(eventType EventType, handler EventHandler) {
	if int(eventType) >= len(e.receivers) { // Check if we have enough room
		// Resize the handler table accordingly
		newRcv := make([][]EventHandler, eventType+1)
		copy(newRcv, e.receivers)
		e.receivers = newRcv
	}

	// Make sure it's a valid event type
	if int(eventType) < len(e.receivers) {
		e.receivers[eventType] = append(e.receivers[eventType], handler)
	}
}

// FireEvent fires an event to all receivers receiving
// Note: if the thread that calls this function communicates with any of the
// callback methods via channels, this function needs to be called as a goroutine
func (e *EventManager) FireEvent(event Event) {
	// No handlers for this event type
	if len(e.receivers) <= int(event.Type()) {
		return
	}

	var wg sync.WaitGroup
	for _, receiver := range e.receivers[event.Type()] {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			receiver.HandleEvent(event) // Why wait for events to get handled?
		}(&wg)
	}
	wg.Wait()
}
