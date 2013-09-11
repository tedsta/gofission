package fission

// Entity ######################################################################

// An abstract object in the scene that contains components
type Entity struct {
	id         int         // Unique id of this entity
	components []Component // Components this entity contains
	typeBits   int         // The combined type bits of the components
}

// AddComponent adds a component to the entity
func (this *Entity) AddComponent(c Component) {
	this.components = append(this.components, c)
	this.typeBits &= c.TypeBits()
}

// GetComponent gets a component attached to this entity by type
// cmpType is the type of component to get
func (this *Entity) GetComponent(cmpType int) Component {
	for _, c := range this.components {
		if c.TypeBits() == cmpType {
			return c
		}
	}
	return nil
}

// GetComponents maps slices of components of the same type to their
// corresponding type bits.
// typeBits specifies which component types to map
func (this *Entity) GetComponents(typeBits int) map[int][]Component {
	cmpMap := make(map[int][]Component)
	for _, c := range this.components {
		if c.TypeBits()&typeBits == c.TypeBits() {
			cmpMap[c.TypeBits()] = append(cmpMap[c.TypeBits()], c)
		}
	}
	return cmpMap
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
