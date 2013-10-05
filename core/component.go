package core

// A type for component type bits
type TypeBits int64

type Component interface {
	Serialize(p *OutPacket)
	Deserialize(p *InPacket)
	TypeBits() TypeBits
}

// This is an internal variable to store the number of registered components
var regCmpntCount uint

type CmpntFactory func() Component

// Store the registered component factories
var componentFactories []CmpntFactory

// Registers a new component type and returns the proper type bits
func RegisterComponent(f CmpntFactory) TypeBits {
	if regCmpntCount >= 64 { // Dat stack overflow doe...
		panic("Cannot have more than 64 component types.")
	}
	defer func() { regCmpntCount++ }()

	componentFactories = append(componentFactories, f)
	return 1 << regCmpntCount
}

// ComponentFactory returns the component factory function for the specified
// type bits
func ComponentFactory(t TypeBits) CmpntFactory {
	if bitIndex(t) < len(componentFactories) {
		return componentFactories[bitIndex(t)]
	}
	return nil
}
