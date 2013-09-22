package event

type Manager struct {
	lCounts  []int        // Listener counts for each event type
	channels []chan Event // Channels for each event type
}

// AddHandler adds an event handler for a particular event type
func (e *Manager) AddListener(eventType Type, l Listener) {
	if int(eventType) >= len(e.lCounts) { // Check if we have enough room
		// Resize arrays accordingly
		newLCounts := make([]int, eventType+1)      // Counts
		newChans := make([]chan Event, eventType+1) // Channels

		copy(newLCounts, e.lCounts)
		copy(newChans, e.channels)

		e.lCounts = newLCounts
		e.channels = newChans

		// Create new count and channel
		e.lCounts[eventType] = 1
		e.channels[eventType] = make(chan Event, 1)
	} else {
		e.lCounts[eventType]++
		e.channels[eventType] = make(chan Event, e.lCounts[eventType])
	}

	go l.Listen(e.channels[eventType])
}

// FireEvent fires an event to all listeners
func (e *Manager) FireEvent(event Event) {
	// No listeners for this event type
	if len(e.lCounts) <= int(event.Type()) {
		return
	}

	for i := 0; i < e.lCounts[event.Type()]; i++ {
		e.channels[event.Type()] <- event
	}
}
