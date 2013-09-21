package rend

import (
	"github.com/tedsta/fission/core"
	"github.com/tedsta/gosfml"
)

type RenderComponent interface {
	core.Component
	Render(t *sf.RenderTarget, states sf.RenderStates)
}
