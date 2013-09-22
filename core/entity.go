package core

// Entity ######################################################################

// An abstract object in the scene that contains components
type Entity struct {
	id         int           // Unique id of this entity
	components [][]Component // Store components in a table based on their type
	typeBits   TypeBits      // The combined type bits of the components
}

// AddComponent adds a component to the entity
func (e *Entity) AddComponent(c Component) {
	bitIndex := bitIndex(c.TypeBits())
	if bitIndex >= len(e.components) { // Check if we have enough room
		// Resize the component table accordingly
		newCmpnts := make([][]Component, bitIndex+1)
		copy(newCmpnts, e.components)
		e.components = newCmpnts
	}

	e.components[bitIndex] = append(e.components[bitIndex], c)
	e.typeBits |= c.TypeBits()
}

// Component gets the first component attached to this entity with the
// specified type
func (e *Entity) Component(typeBits TypeBits) Component {
	// No space in table for the component - that means it doesn't exist
	if bitIndex(typeBits) >= len(e.components) ||
		len(e.components[bitIndex(typeBits)]) == 0 {
		return nil
	}
	return e.components[bitIndex(typeBits)][0]
}

// Components returns a slice of all the components in this entity with the
// specified type
func (e *Entity) Components(typeBits TypeBits) []Component {
	return e.components[bitIndex(typeBits)]
}

// Id returns the id of the entity
func (e *Entity) Id() int {
	return e.id
}

// entityEvent #################################################################

// An event type for entity events.
type entityEvent struct {
	eventType EventType // There can be more than one type of entity event
	Ent       *Entity   // The entity this event is referring to
}

func (e *entityEvent) Type() EventType {
	return e.eventType
}

// util ########################################################################

// bitIndex returns the index of a single set bit in an integer
func bitIndex(val TypeBits) int {
	index := 0
	shift := func(i *TypeBits) TypeBits { *i >>= 1; return *i }
	for shift(&val) > 0 {
		index++
	}
	return index
}
