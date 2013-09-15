package input

import (
	"fission"
	"github.com/tedsta/go-sfml"
)

var IntentComponentType = fission.NextComponentType()

type IntentComponent struct {
	keyMap  [sfml.KeyCount]string // Maps key codes to intent names
	intents map[string]bool       // Maps intent names to their state
}

func NewIntentComponent() *IntentComponent {
	intents := make(map[string]bool)
	return &IntentComponent{[sfml.KeyCount]string{}, intents}
}

func (i *IntentComponent) Serialize() {
}

func (i *IntentComponent) Deserialize() {

}

func (i *IntentComponent) MapKeyToIntent(key sfml.KeyCode, intent string) {
	i.keyMap[uint(key)] = intent
}

func (i *IntentComponent) IsIntentActive(intent string) bool {
	return i.intents[intent]
}

func (i *IntentComponent) OnKeyPressed(key sfml.KeyCode) {
	if uint(key) < uint(sfml.KeyCount) {
		i.intents[i.keyMap[uint(key)]] = true
	}
}

func (i *IntentComponent) OnKeyReleased(key sfml.KeyCode) {
	if uint(key) < uint(sfml.KeyCount) {
		i.intents[i.keyMap[uint(key)]] = false
	}
}

func (i *IntentComponent) TypeBits() fission.TypeBits {
	return IntentComponentType
}
