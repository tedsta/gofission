package rnd

import (
	"github.com/tedsta/fission/core"
	"github.com/tedsta/gosfml"
)

/*type RenderComponent interface {
	core.Component
	Render(t *sf.RenderTarget, states sf.RenderStates)
}*/

var RenderComponentType = core.NextComponentType()

type RenderComponent struct {
	serialize   func(*core.OutPacket)
	deserialize func(*core.InPacket)
	Render      func(*sf.RenderTarget, sf.RenderStates) // Render function
}

func NewRenderComponent(serialize func(p *core.OutPacket), deserialize func(p *core.InPacket),
	render func(*sf.RenderTarget, sf.RenderStates)) *RenderComponent {

	return &RenderComponent{serialize, deserialize, render}
}

func (r *RenderComponent) Serialize(p *core.OutPacket) {
	r.serialize(p)
}

func (r *RenderComponent) Deserialize(p *core.InPacket) {
	r.deserialize(p)
}

func (r *RenderComponent) TypeBits() core.TypeBits {
	return RenderComponentType
}
