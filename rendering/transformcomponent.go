package rnd

import (
	"github.com/tedsta/fission/core"
	"github.com/tedsta/gosfml"
)

// The type bits for TransformComponent
var TransformComponentType core.TypeBits

type TransformComponent struct {
	T *sf.Transformable
}

func TransformComponentFactory() core.Component {
	return NewTransformComponent()
}

func NewTransformComponent() *TransformComponent {
	return &TransformComponent{sf.NewTransformable()}
}

func (t *TransformComponent) Serialize(p *core.OutPacket) {
	p.Write(t.T.Origin().X, t.T.Origin().Y)
	p.Write(t.T.Position().X, t.T.Position().Y)
	p.Write(t.T.Rotation())
	p.Write(t.T.Scale().X, t.T.Scale().Y)
}

func (t *TransformComponent) Deserialize(p *core.InPacket) {
	var origin, pos, scale sf.Vector2
	var rot float32
	p.Read(&origin.X, &origin.Y)
	p.Read(&pos.X, &pos.Y)
	p.Read(&rot)
	p.Read(&scale.X, &scale.Y)

	t.T.SetOrigin(origin)
	t.T.SetPosition(pos)
	t.T.SetRotation(rot)
	t.T.SetScale(scale)
}

func (t *TransformComponent) TypeBits() core.TypeBits {
	return TransformComponentType
}
