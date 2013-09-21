package input

import (
	"fmt"
	glfw "github.com/go-gl/glfw3"
	"github.com/tedsta/fission/core"
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
	eventManager *core.EventManager
	window       *glfw.Window
	keyStates    [glfw.KeyLast]BtnState
	typeBits     core.TypeBits // Extra type bits for custom input components
}

func NewInputSystem(w *glfw.Window, e *core.EventManager, typeBits core.TypeBits) *InputSystem {
	i := &InputSystem{}

	// Set the input callbacks
	w.SetFramebufferSizeCallback((i).onResize)
	w.SetMouseButtonCallback((i).onMouseBtn)
	//w.SetMouseWheelCallback((i).onMouseWheel)
	w.SetKeyCallback((i).onKey)
	w.SetCharacterCallback((i).onChar)

	i.window = w
	i.eventManager = e
	i.keyStates = [glfw.KeyLast]BtnState{}
	i.typeBits = typeBits

	return i
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

func (i *InputSystem) onResize(wnd *glfw.Window, w, h int) {
	fmt.Printf("resized: %dx%d\n", w, h)
}

func (i *InputSystem) onMouseBtn(w *glfw.Window, btn glfw.MouseButton, act glfw.Action, mod glfw.ModifierKey) {
	fmt.Printf("mouse button: %d, %d\n", btn, act)
}

func (i *InputSystem) onMouseWheel(w *glfw.Window, delta int) {
	fmt.Printf("mouse wheel: %d\n", delta)
}

func (i *InputSystem) onKey(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	e := &KeyEvent{key, scancode, action, mods}
	i.eventManager.FireEvent(e)
}

func (i *InputSystem) onChar(w *glfw.Window, key uint) {
	fmt.Printf("char: %d\n", key)
}
