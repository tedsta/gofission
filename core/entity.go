package core

import (
	"github.com/tedsta/fission/core/event"
)

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
	cmpnts := []Component{}
	indices := bitIndices(typeBits)
	for _, i := range indices {
		cmpnts = append(cmpnts, e.components[i]...)
	}
	return cmpnts
}

// IsType returns true if the entity's type bits contain the type bits supplied
// Example:
// Sprite := TransformComponentType | SpriteComponentType
// if e.IsType(Sprite) { /*e is a sprite*/ }
func (e *Entity) IsType(t TypeBits) bool {
	return e.typeBits&t == t
}

// Id returns the id of the entity
func (e *Entity) Id() int {
	return e.id
}

func (e *Entity) TypeBits() TypeBits {
	return e.typeBits
}

// entityEvent #################################################################

// An event type for entity events.
type entityEvent struct {
	eventType event.Type // There can be more than one type of entity event
	Ent       *Entity    // The entity this event is referring to
}

func (e *entityEvent) Type() event.Type {
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

// bitIndices returns a slice of all the set bit indices
func bitIndices(val TypeBits) []uint8 {
	bits := []uint8{}
	shifter := TypeBits(1)
	for i := 0; i < 64; i++ {
		if val&shifter != 0 {
			bits = append(bits, uint8(i))
		}
		shifter <<= 1
	}
	return bits
}
