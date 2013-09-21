package input

import (
	glfw "github.com/go-gl/glfw3"
	"github.com/tedsta/fission/core"
)

// KeyEvent ####################################################################
var KeyEventType = core.NextEventId()

type KeyEvent struct {
	Key      glfw.Key
	Scancode int
	Action   glfw.Action
	Mods     glfw.ModifierKey
}

func (k *KeyEvent) Type() core.EventType {
	return KeyEventType
}
