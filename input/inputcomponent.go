package input

import (
	"fission"
	"github.com/tedsta/go-sfml"
)

type InputComponent interface {
	fission.Component
	OnKeyPressed(key sfml.KeyCode)
	OnKeyReleased(key sfml.KeyCode)
}
