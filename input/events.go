package input

import (
	"github.com/tedsta/fission/core/event"
)

// KeyEvent ####################################################################
var KeyEventType = event.NextId()

type KeyEvent struct {
	Key      Key
	Scancode int
	Action   Action
	Mods     ModifierKey
}

func (k *KeyEvent) Type() event.Type {
	return KeyEventType
}
