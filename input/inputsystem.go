package input

import (
	"fission"
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
	window    *sfml.RenderWindow
	keyStates [sfml.KeyCount]BtnState
	events    []interface{}
	typeBits  fission.TypeBits // Extra type bits for custom input components
}

func NewInputSystem(win *sfml.RenderWindow, typeBits fission.TypeBits) *InputSystem {
	return &InputSystem{win, [sfml.KeyCount]BtnState{}, nil, typeBits}
}

func (i *InputSystem) Begin(dt float32) {
	// Initialize the slice of events
	i.events = make([]interface{}, 0, 1)

	for {
		e, ok := i.window.PollEvent()
		if !ok {
			break
		}

		i.events = append(i.events, e)
	}
}

func (i *InputSystem) ProcessEntity(e *fission.Entity, dt float32) {
	cmpnts := e.Components(i.TypeBits()) // Grab all of the input components

	// Convert the components to input components
	inputCmpnts := make([]InputComponent, len(cmpnts))
	for i, cmpnt := range cmpnts {
		inputCmpnts[i] = cmpnt.(InputComponent)
	}

	// Send the events to all of the input components
	for _, event := range i.events {
		for _, cmpnt := range inputCmpnts {
			switch event.(type) {
			// Keyboard event
			case sfml.KeyEvent:
				evt := event.(sfml.KeyEvent)
				if evt.Type == sfml.EvtKeyPressed {
					cmpnt.OnKeyPressed(evt.Code())
				} else if evt.Type == sfml.EvtKeyReleased {
					cmpnt.OnKeyReleased(evt.Code())
				}
			}
		}
	}
}

func (i *InputSystem) End(dt float32) {
}

func (i *InputSystem) TypeBits() fission.TypeBits {
	return IntentComponentType | i.typeBits
}
