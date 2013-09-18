package input

import (
	"fission/core"
	glfw "github.com/go-gl/glfw3"
)

var IntentComponentType = core.NextComponentType()

type IntentComponent struct {
	keyMap  [glfw.KeyLast]string // Maps key codes to intent names
	intents map[string]bool      // Maps intent names to their state
}

func NewIntentComponent() *IntentComponent {
	intents := make(map[string]bool)
	return &IntentComponent{[glfw.KeyLast]string{}, intents}
}

func (i *IntentComponent) Serialize() {
}

func (i *IntentComponent) Deserialize() {

}

func (i *IntentComponent) TypeBits() core.TypeBits {
	return IntentComponentType
}

func (i *IntentComponent) MapKeyToIntent(key glfw.Key, intent string) {
	i.keyMap[uint(key)] = intent
}

func (i *IntentComponent) IsIntentActive(intent string) bool {
	return i.intents[intent]
}

func (i *IntentComponent) OnKeyPressed(key glfw.Key) {
	if key < glfw.KeyLast {
		i.intents[i.keyMap[uint(key)]] = true
	}
}

func (i *IntentComponent) OnKeyReleased(key glfw.Key) {
	if key < glfw.KeyLast {
		i.intents[i.keyMap[uint(key)]] = false
	}
}
