package rend

import (
	"fission"
	"github.com/tedsta/go-sfml"
)

type RenderComponent interface {
	fission.Component
	Render(win sfml.RenderWindow, pos sfml.Vector2f, rot, scale float32)
}
