package rend

import (
	"github.com/tedsta/fission/core"
)

// The type bits for TransformComponent
var TransformComponentType = core.NextComponentType()

type TransformComponent struct {
	Pos   Vector2
	Rot   float32
	Scale float32
}

func NewTransformComponent() *TransformComponent {
	return &TransformComponent{Vector2{}, 0, 1.0}
}

func (t *TransformComponent) Serialize() {
}

func (t *TransformComponent) Deserialize() {
}

func (t *TransformComponent) TypeBits() core.TypeBits {
	return TransformComponentType
}
