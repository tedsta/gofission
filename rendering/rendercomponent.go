package rendering

import (
	"fission"
	"github.com/tedsta/go-sfml"
)

type RenderComponent interface {
	fission.Component
	Render(window sfml.RenderWindow, pos sfml.Vector2f)
}
