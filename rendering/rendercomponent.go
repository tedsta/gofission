package rend

import (
	"github.com/tedsta/fission/core"
)

type RenderComponent interface {
	core.Component
	Render(t *RenderTarget, states RenderStates, pos Vector2, rot, scale float32)
}
