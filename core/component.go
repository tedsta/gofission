package core

// A type for component type bits
type TypeBits int64

// This is an internal variable to store the number of registered components
var regCmpntCount uint = 0

// Registers a new component type and returns the proper type bits
func NextComponentType() TypeBits {
	if regCmpntCount >= 64 { // Dat stack overflow doe...
		panic("Cannot have more than 64 component types.")
	}

	defer func() { regCmpntCount++ }()
	return 1 << regCmpntCount
}

type Component interface {
	Serialize(p *Packet)
	Deserialize(p *Packet)
	TypeBits() TypeBits
}
