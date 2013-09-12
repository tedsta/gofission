package fission

import (
	"github.com/tedsta/go-sfml"
)

var TransformComponentType = NextComponentType()

type TransformComponent struct {
	Position sfml.Vector2f
	Rotation float32
	Scale    float32
}

func NewTransformComponent() *TransformComponent {
	return &TransformComponent{sfml.NewVector2f(0, 0), 0, 1.0}
}

func (this *TransformComponent) Serialize() {
}

func (this *TransformComponent) Deserialize() {
}

func (this *TransformComponent) TypeBits() int {
	return TransformComponentType
}
