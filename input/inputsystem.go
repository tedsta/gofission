package input

import (
	glfw "github.com/go-gl/glfw3"
	"github.com/tedsta/fission/core"
	"github.com/tedsta/fission/core/event"
)

type InputSystem struct {
	eventManager *event.Manager
	window       *glfw.Window
}

func NewInputSystem(w *glfw.Window, e *event.Manager) *InputSystem {
	i := &InputSystem{}

	// Set the input callbacks
	w.SetMouseButtonCallback(i.onMouseBtn)
	//w.SetMouseWheelCallback((i).onMouseWheel)
	w.SetCursorPositionCallback(i.onMouseMove)
	w.SetKeyCallback(i.onKey)
	w.SetCharacterCallback(i.onChar)

	i.window = w
	i.eventManager = e

	return i
}

func (i *InputSystem) Begin(dt float32) {
	glfw.PollEvents()
}

func (i *InputSystem) ProcessEntity(e *core.Entity, dt float32) {
}

func (i *InputSystem) End(dt float32) {
}

func (i *InputSystem) TypeBits() (core.TypeBits, core.TypeBits) {
	return 0, 0
}

// Callbacks ###################################################################

func (i *InputSystem) onResize(wnd *glfw.Window, w, h int) {
	//fmt.Printf("resized: %dx%d\n", w, h)
}

func (i *InputSystem) onMouseBtn(w *glfw.Window, btn glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	e := &MouseBtnEvent{MouseButton(btn), Action(action), ModifierKey(mods)}
	i.eventManager.FireEvent(e)
}

func (i *InputSystem) onMouseWheel(w *glfw.Window, delta int) {
	//fmt.Printf("mouse wheel: %d\n", delta)
}

func (i *InputSystem) onMouseMove(w *glfw.Window, xpos float64, ypos float64) {
	e := &MouseMoveEvent{int(xpos), int(ypos)}
	i.eventManager.FireEvent(e)
}

func (i *InputSystem) onKey(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	e := &KeyEvent{Key(key), scancode, Action(action), ModifierKey(mods)}
	i.eventManager.FireEvent(e)
}

func (i *InputSystem) onChar(w *glfw.Window, key uint) {
	//fmt.Printf("char: %d\n", key)
}

// init ##########################################################

func init() {
	IntentComponentType = core.RegisterComponent(IntentComponentFactory)
}
