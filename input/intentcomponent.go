package input

import (
	glfw "github.com/go-gl/glfw3"
	"github.com/tedsta/fission/core"
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

func (i *IntentComponent) IntentActive(intent string) bool {
	return i.intents[intent]
}

func (i *IntentComponent) HandleEvent(e core.Event) {
	switch e.Type() {
	case KeyEventType:
		ke := e.(*KeyEvent)
		if ke.Action == glfw.Press {
			i.intents[i.keyMap[uint(ke.Key)]] = true
		} else if ke.Action == glfw.Release {
			i.intents[i.keyMap[uint(ke.Key)]] = false
		}
	}
}
