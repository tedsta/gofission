package input

import (
	"github.com/tedsta/fission/core"
	"github.com/tedsta/fission/core/event"
)

var IntentComponentType = core.NextComponentType()

type IntentComponent struct {
	intent *IntentMapper
}

func NewIntentComponent() *IntentComponent {
	intent := NewIntentMapper()
	return &IntentComponent{intent}
}

func (i *IntentComponent) TypeBits() core.TypeBits {
	return IntentComponentType
}

func (i *IntentComponent) MapKeyToIntent(key Key, intent string) {
	i.intent.MapKeyToIntent(key, intent)
}

func (i *IntentComponent) IntentActive(intent string) bool {
	return i.intent.IntentActive(intent)
}

// IntentMapper ################################################################

type IntentMapper struct {
	keyMap  [KeyLast]string // Maps key codes to intent names
	intents map[string]bool // Maps intent names to their state
}

func NewIntentMapper() *IntentMapper {
	intents := make(map[string]bool)
	return &IntentMapper{[KeyLast]string{}, intents}
}

func (i *IntentMapper) MapKeyToIntent(key Key, intent string) {
	i.keyMap[uint(key)] = intent
}

func (i *IntentMapper) IntentActive(intent string) bool {
	return i.intents[intent]
}

func (i *IntentMapper) HandleEvent(e event.Event) {
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
