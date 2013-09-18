package rend

import (
	"fission/core"
)

type RenderComponent interface {
	core.Component
	Render(pos *core.Vector2, rot, scale float32)
}
