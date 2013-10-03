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

// MouseBtnEvent ####################################################################
var MouseBtnEventType = event.NextId()

type MouseBtnEvent struct {
	Btn    MouseButton
	Action Action
	Mods   ModifierKey
}

func (m *MouseBtnEvent) Type() event.Type {
	return MouseBtnEventType
}

// MouseMoveEvent ####################################################################
var MouseMoveEventType = event.NextId()

type MouseMoveEvent struct {
	X, Y int
}

func (m *MouseMoveEvent) Type() event.Type {
	return MouseMoveEventType
}
