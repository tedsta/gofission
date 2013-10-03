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
	serialize   func(*core.Packet)
	deserialize func(*core.Packet)
	Render      func(*sf.RenderTarget, sf.RenderStates) // Render function
}

func NewRenderComponent(serialize, deserialize func(p *core.Packet),
	render func(*sf.RenderTarget, sf.RenderStates)) *RenderComponent {

	return &RenderComponent{serialize, deserialize, render}
}

func (r *RenderComponent) Serialize(p *core.Packet) {
	r.serialize(p)
}

func (r *RenderComponent) Deserialize(p *core.Packet) {
	r.deserialize(p)
}

func (r *RenderComponent) TypeBits() core.TypeBits {
	return RenderComponentType
}
