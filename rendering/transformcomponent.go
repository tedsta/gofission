package rend

import (
	"github.com/tedsta/fission/core"
)

// The type bits for TransformComponent
var TransformComponentType = core.NextComponentType()

type TransformComponent struct {
	T *Transformable
}

func NewTransformComponent() *TransformComponent {
	return &TransformComponent{NewTransformable()}
}

func (t *TransformComponent) Serialize() {
}

func (t *TransformComponent) Deserialize() {
}

func (t *TransformComponent) TypeBits() core.TypeBits {
	return TransformComponentType
}
