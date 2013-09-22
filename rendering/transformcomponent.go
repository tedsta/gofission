package rend

import (
	"github.com/tedsta/fission/core"
	"github.com/tedsta/gosfml"
)

// The type bits for TransformComponent
var TransformComponentType = core.NextComponentType()

type TransformComponent struct {
	T *sf.Transformable
}

func NewTransformComponent() *TransformComponent {
	return &TransformComponent{sf.NewTransformable()}
}

func (t *TransformComponent) TypeBits() core.TypeBits {
	return TransformComponentType
}
