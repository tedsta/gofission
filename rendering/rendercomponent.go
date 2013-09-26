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
	Render func(t *sf.RenderTarget, states sf.RenderStates) // Render function
}

func (r *RenderComponent) TypeBits() core.TypeBits {
	return RenderComponentType
}
