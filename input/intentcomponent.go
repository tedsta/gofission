package input

import (
	"github.com/tedsta/fission/core"
	"github.com/tedsta/fission/core/event"
)

var IntentComponentType core.TypeBits

type IntentComponent struct {
	intent *IntentMapper
}

func IntentComponentFactory() core.Component {
	return &IntentComponent{}
}

func NewIntentComponent(evt *event.Manager) *IntentComponent {
	intent := NewIntentMapper(evt)
	return &IntentComponent{intent}
}

func (i *IntentComponent) Serialize(p *core.OutPacket) {
}

func (i *IntentComponent) Deserialize(p *core.InPacket) {
}

func (i *IntentComponent) TypeBits() core.TypeBits {
	return IntentComponentType
}

func (i *IntentComponent) MapKeyToIntent(key Key, state BtnState, intent string) {
	i.intent.MapKeyToIntent(key, state, intent)
}

func (i *IntentComponent) MapMouseBtnToIntent(btn MouseButton, state BtnState, intent string) {
	i.intent.MapMouseBtnToIntent(btn, state, intent)
}

func (i *IntentComponent) IntentActive(intent string) bool {
	return i.intent.IntentActive(intent)
}

// IntentMapper ################################################################

type BtnState uint8

const (
	Up BtnState = iota
	Down
	Released
	Pressed
)

type input struct {
	inputType InputMode
	value     int
	state     BtnState
}

type IntentMapper struct {
	inputMap map[input]string // Maps input to intent names
	intents  map[string]*bool // Maps intent names to their state

	keyStates      [KeyLast]BtnState
	mouseStates    [MouseButtonLast]BtnState
	mouseX, mouseY int
}

func NewIntentMapper(evt *event.Manager) *IntentMapper {
	intents := make(map[string]*bool)

	i := &IntentMapper{inputMap: map[input]string{}, intents: intents}
	evt.AddHandler(KeyEventType, i)
	evt.AddHandler(MouseBtnEventType, i)
	evt.AddHandler(MouseMoveEventType, i)
	return i
}

// Call this when done processing input for the current frame
func (i *IntentMapper) Finish() {
	for _, v := range i.intents {
		*v = false
	}

	for k := 0; k < int(KeyLast); k++ {
		if i.keyStates[k] == Pressed {
			i.keyStates[k] = Down
		} else if i.keyStates[k] == Released {
			i.keyStates[k] = Up
		}

		in := input{StickyKeys, k, i.keyStates[k]}

		if _, ok := i.inputMap[in]; ok {
			*i.intents[i.inputMap[in]] = true
		}
	}

	for m := 0; m < int(MouseButtonLast); m++ {
		if i.mouseStates[m] == Pressed {
			i.mouseStates[m] = Down
		} else if i.mouseStates[m] == Released {
			i.mouseStates[m] = Up
		}

		in := input{StickyMouseButtons, m, i.mouseStates[m]}

		if _, ok := i.inputMap[in]; ok {
			*i.intents[i.inputMap[in]] = true
		}
	}
}

func (i *IntentMapper) MapKeyToIntent(key Key, state BtnState, intent string) {
	in := input{StickyKeys, int(key), state}
	i.inputMap[in] = intent
	i.intents[intent] = new(bool)
}

func (i *IntentMapper) MapMouseBtnToIntent(btn MouseButton, state BtnState, intent string) {
	in := input{StickyMouseButtons, int(btn), state}
	i.inputMap[in] = intent
	i.intents[intent] = new(bool)
}

func (i *IntentMapper) IntentActive(intent string) bool {
	return *i.intents[intent]
}

func (i *IntentMapper) MousePos() (x, y int) {
	return i.mouseX, i.mouseY
}

func (i *IntentMapper) HandleEvent(e event.Event) {
	switch e.Type() {
	case KeyEventType:
		ke := e.(*KeyEvent)

		if ke.Action == Press {
			i.keyStates[int(ke.Key)] = Pressed

			in := input{StickyKeys, int(ke.Key), Pressed}
			if _, ok := i.inputMap[in]; ok {
				*i.intents[i.inputMap[in]] = true
			}
		} else if ke.Action == Release {
			i.keyStates[int(ke.Key)] = Released

			in := input{StickyKeys, int(ke.Key), Released}
			if _, ok := i.inputMap[in]; ok {
				*i.intents[i.inputMap[in]] = true
			}
		}
	case MouseBtnEventType:
		me := e.(*MouseBtnEvent)

		if me.Action == Press {
			i.mouseStates[int(me.Btn)] = Pressed

			in := input{StickyMouseButtons, int(me.Btn), Pressed}
			if _, ok := i.inputMap[in]; ok {
				*i.intents[i.inputMap[in]] = true
			}
		} else if me.Action == Release {
			i.mouseStates[int(me.Btn)] = Released

			in := input{StickyMouseButtons, int(me.Btn), Released}
			if _, ok := i.inputMap[in]; ok {
				*i.intents[i.inputMap[in]] = true
			}
		}
	case MouseMoveEventType:
		me := e.(*MouseMoveEvent)
		i.mouseX, i.mouseY = me.X, me.Y
	}
}
