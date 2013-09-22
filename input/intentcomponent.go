package input

import (
	"github.com/tedsta/fission/core"
)

var IntentComponentType = core.NextComponentType()

type IntentComponent struct {
	keyMap  [KeyLast]string // Maps key codes to intent names
	intents map[string]bool // Maps intent names to their state
}

func NewIntentComponent() *IntentComponent {
	intents := make(map[string]bool)
	return &IntentComponent{[KeyLast]string{}, intents}
}

func (i *IntentComponent) TypeBits() core.TypeBits {
	return IntentComponentType
}

func (i *IntentComponent) MapKeyToIntent(key Key, intent string) {
	i.keyMap[uint(key)] = intent
}

func (i *IntentComponent) IntentActive(intent string) bool {
	return i.intents[intent]
}

func (i *IntentComponent) HandleEvent(e core.Event) {
	switch e.Type() {
	case KeyEventType:
		ke := e.(*KeyEvent)
		if ke.Action == Press {
			i.intents[i.keyMap[uint(ke.Key)]] = true
		} else if ke.Action == Release {
			i.intents[i.keyMap[uint(ke.Key)]] = false
		}
	}
}
