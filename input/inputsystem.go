package fission

import (
	"github.com/tedsta/go-sfml"
)

// Type for button states
type BtnState byte

// Enumerate button states
const (
	BtnUp BtnState = iota
	BtnReleased
	BtnDown
	BtnPressed
)

type InputSystem struct {
	keyStates [sfml.KeyCount]BtnState
}
