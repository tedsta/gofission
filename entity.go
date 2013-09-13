package fission

// Entity ######################################################################

// An abstract object in the scene that contains components
type Entity struct {
	id         int           // Unique id of this entity
	components [][]Component // Store components in a table based on their type
	typeBits   TypeBits      // The combined type bits of the components
}

// AddComponent adds a component to the entity
func (this *Entity) AddComponent(c Component) {
	bitIndex := bitIndex(c.TypeBits())
	if bitIndex >= len(this.components) { // Check if we have enough room
		// Resize the component table accordingly
		this.components = append(this.components, make([]Component,
			bitIndex-len(this.components)))
	}

	this.components[bitIndex] = append(this.components[bitIndex], c)
	this.typeBits |= c.TypeBits()
}

// Component gets the first component attached to this entity with the
// specified type
func (this *Entity) Component(typeBits TypeBits) Component {
	// No space in table for the component - that means it doesn't exist
	if bitIndex(typeBits) >= len(this.components) ||
		len(this.components[bitIndex(typeBits)]) == 0 {
		return nil
	}
	return this.components[bitIndex(typeBits)][0]
}

// Components returns a slice of all the components in this entity with the
// specified type
func (this *Entity) Components(typeBits TypeBits) []Component {
	return this.components[bitIndex(typeBits)]
}

// Serialize serializes the entity into a packet
func (this *Entity) Serialize() {
	// TODO: Implement this
}

// Deserialize deserializes the entity from a packet
func (this *Entity) Deserialize() {
	// TODO: Implement this
}

// Id returns the id of the entity
func (this *Entity) Id() int {
	return this.id
}

// entityEvent #################################################################

// An event type for entity events.
type entityEvent struct {
	eventType int     // There can be more than one type of entity event
	Ent       *Entity // The entity this event is referring to
}

func (this *entityEvent) Type() int {
	return this.eventType
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
