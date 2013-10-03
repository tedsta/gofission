package rnd

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

func (t *TransformComponent) Serialize(p *core.OutPacket) {
	p.Write(t.T.Origin())
	p.Write(t.T.Position())
	p.Write(t.T.Rotation())
	p.Write(t.T.Scale())
}

func (t *TransformComponent) Deserialize(p *core.InPacket) {
	var origin, pos, scale sf.Vector2
	var rot float32
	p.Read(&origin)
	p.Read(&pos)
	p.Read(&rot)
	p.Read(&scale)

	t.T.SetOrigin(origin)
	t.T.SetPosition(pos)
	t.T.SetRotation(rot)
	t.T.SetScale(scale)
}

func (t *TransformComponent) TypeBits() core.TypeBits {
	return TransformComponentType
}
