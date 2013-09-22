package input

import (
	"github.com/tedsta/fission/core"
)

// KeyEvent ####################################################################
var KeyEventType = core.NextEventId()

type KeyEvent struct {
	Key      Key
	Scancode int
	Action   Action
	Mods     ModifierKey
}

func (k *KeyEvent) Type() core.EventType {
	return KeyEventType
}
