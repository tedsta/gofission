package input

import (
	"fission/core"
	"fmt"
	glfw "github.com/go-gl/glfw3"
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
	window    *glfw.Window
	keyStates [glfw.KeyLast]BtnState
	typeBits  core.TypeBits // Extra type bits for custom input components
}

func NewInputSystem(w *glfw.Window, typeBits core.TypeBits) *InputSystem {
	// Set the input callbacks
	w.SetFramebufferSizeCallback(onResize)
	w.SetMouseButtonCallback(onMouseBtn)
	//w.SetMouseWheelCallback(onMouseWheel)
	w.SetKeyCallback(onKey)
	w.SetCharacterCallback(onChar)
	return &InputSystem{w, [glfw.KeyLast]BtnState{}, typeBits}
}

func (i *InputSystem) Begin(dt float32) {
	glfw.PollEvents()
}

func (i *InputSystem) ProcessEntity(e *core.Entity, dt float32) {
	cmpnts := e.Components(i.TypeBits()) // Grab all of the input components

	// Convert the components to input components
	inputCmpnts := make([]InputComponent, len(cmpnts))
	for i, cmpnt := range cmpnts {
		inputCmpnts[i] = cmpnt.(InputComponent)
	}

	/*for _, cmpnt := range inputCmpnts {
	}*/
}

func (i *InputSystem) End(dt float32) {
}

func (i *InputSystem) TypeBits() core.TypeBits {
	return IntentComponentType | i.typeBits
}

// Callbacks ###################################################################

func onResize(wnd *glfw.Window, w, h int) {
	fmt.Printf("resized: %dx%d\n", w, h)
}

func onMouseBtn(w *glfw.Window, btn glfw.MouseButton, act glfw.Action, mod glfw.ModifierKey) {
	fmt.Printf("mouse button: %d, %d\n", btn, act)
}

func onMouseWheel(w *glfw.Window, delta int) {
	fmt.Printf("mouse wheel: %d\n", delta)
}

func onKey(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	fmt.Printf("key: %d, %d, %d\n", key, action, mods)
}

func onChar(w *glfw.Window, key uint) {
	fmt.Printf("char: %d\n", key)
}
