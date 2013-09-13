package fission

import (
	"github.com/tedsta/go-sfml"
)

// The type bits for TransformComponent
var TransformComponentType = NextComponentType()

type TransformComponent struct {
	Pos   sfml.Vector2f
	Rot   float32
	Scale float32
}

func NewTransformComponent() *TransformComponent {
	return &TransformComponent{sfml.NewVector2f(0, 0), 0, 1.0}
}

func (t *TransformComponent) Serialize() {
}

func (t *TransformComponent) Deserialize() {
}

func (t *TransformComponent) TypeBits() TypeBits {
	return TransformComponentType
}
