package net

import (
	"github.com/tedsta/fission/core"
	"github.com/tedsta/fission/core/event"
	"github.com/tedsta/fission/input"
)

var IntentComponentType core.TypeBits

type IntentComponent struct {
	intent *IntentMapper
}

func IntentComponentFactory() core.Component {
	return &IntentComponent{}
}

func NewIntentComponent(evt *event.Manager, conn *Connection) *IntentComponent {
	intent := NewIntentMapper(evt, conn)
	return &IntentComponent{intent}
}

func (i *IntentComponent) Serialize(p *core.OutPacket) {
	p.Write(i.intent.hndId)
}

func (i *IntentComponent) Deserialize(p *core.InPacket) {
	p.Read(&i.intent.hndId)
}

func (i *IntentComponent) TypeBits() core.TypeBits {
	return IntentComponentType
}

func (i *IntentComponent) MapKeyToIntent(key input.Key, state input.BtnState, intent string) {
	i.intent.MapKeyToIntent(key, state, intent)
}

func (i *IntentComponent) MapMouseBtnToIntent(btn input.MouseButton, state input.BtnState, intent string) {
	i.intent.MapMouseBtnToIntent(btn, state, intent)
}

func (i *IntentComponent) Finish() {
	i.intent.Finish()
}

func (i *IntentComponent) IntentActive(intent string) bool {
	return i.intent.IntentActive(intent)
}

func (i *IntentComponent) MousePos() (x, y int) {
	return i.intent.mouseX, i.intent.mouseY
}

// IntentMapper ################################################################

// Packet types
const (
	KeyEvent int = iota
	MouseBtnEvent
	MouseMoveEvent
)

type action struct {
	inputType input.InputMode
	value     int
	state     input.BtnState
}

type IntentMapper struct {
	conn     *Connection
	hndId    int
	inputMap map[action]string // Maps input to intent names
	intents  map[string]*bool  // Maps intent names to their state

	keyStates      [input.KeyLast]input.BtnState
	mouseStates    [input.MouseButtonLast]input.BtnState
	mouseX, mouseY int
}

func NewIntentMapper(evt *event.Manager, conn *Connection) *IntentMapper {
	intents := make(map[string]*bool)

	i := &IntentMapper{conn: conn, inputMap: map[action]string{}, intents: intents}
	if conn.Type() == Client {
		evt.AddHandler(input.KeyEventType, i)
		evt.AddHandler(input.MouseBtnEventType, i)
		evt.AddHandler(input.MouseMoveEventType, i)
	} else {
		i.hndId = conn.RegisterHandlerAuto(i.HandlePacket)
	}
	return i
}

// Call this when done processing input for the current frame
func (i *IntentMapper) Finish() {
	for _, v := range i.intents {
		*v = false
	}

	for k := 0; k < int(input.KeyLast); k++ {
		if i.keyStates[k] == input.Pressed {
			i.keyStates[k] = input.Down
		} else if i.keyStates[k] == input.Released {
			i.keyStates[k] = input.Up
		}

		in := action{input.StickyKeys, k, i.keyStates[k]}

		// Send it across the network
		packet := core.NewOutPacket(nil)
		packet.Write(KeyEvent, input.StickyKeys, k, i.keyStates[k])
		i.conn.Send(packet, i.hndId, 0, 0, false)

		if _, ok := i.inputMap[in]; ok {
			*i.intents[i.inputMap[in]] = true
		}
	}

	for m := 0; m < int(input.MouseButtonLast); m++ {
		if i.mouseStates[m] == input.Pressed {
			i.mouseStates[m] = input.Down
		} else if i.mouseStates[m] == input.Released {
			i.mouseStates[m] = input.Up
		}

		in := action{input.StickyMouseButtons, m, i.mouseStates[m]}

		// Send it across the network
		packet := core.NewOutPacket(nil)
		packet.Write(KeyEvent, input.StickyKeys, m, i.keyStates[m])
		i.conn.Send(packet, i.hndId, 0, 0, false)

		if _, ok := i.inputMap[in]; ok {
			*i.intents[i.inputMap[in]] = true
		}
	}
}

func (i *IntentMapper) MapKeyToIntent(key input.Key, state input.BtnState, intent string) {
	in := action{input.StickyKeys, int(key), state}
	i.inputMap[in] = intent
	i.intents[intent] = new(bool)
}

func (i *IntentMapper) MapMouseBtnToIntent(btn input.MouseButton, state input.BtnState, intent string) {
	in := action{input.StickyMouseButtons, int(btn), state}
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
	case input.KeyEventType:
		ke := e.(*input.KeyEvent)

		if ke.Action == input.Press {
			i.keyStates[int(ke.Key)] = input.Pressed

			in := action{input.StickyKeys, int(ke.Key), input.Pressed}
			if _, ok := i.inputMap[in]; ok {
				*i.intents[i.inputMap[in]] = true
			}
		} else if ke.Action == input.Release {
			i.keyStates[int(ke.Key)] = input.Released

			in := action{input.StickyKeys, int(ke.Key), input.Released}
			if _, ok := i.inputMap[in]; ok {
				*i.intents[i.inputMap[in]] = true
			}
		}
	case input.MouseBtnEventType:
		me := e.(*input.MouseBtnEvent)

		if me.Action == input.Press {
			i.mouseStates[int(me.Btn)] = input.Pressed

			in := action{input.StickyMouseButtons, int(me.Btn), input.Pressed}
			if _, ok := i.inputMap[in]; ok {
				*i.intents[i.inputMap[in]] = true
			}
		} else if me.Action == input.Release {
			i.mouseStates[int(me.Btn)] = input.Released

			in := action{input.StickyMouseButtons, int(me.Btn), input.Released}
			if _, ok := i.inputMap[in]; ok {
				*i.intents[i.inputMap[in]] = true
			}
		}
	case input.MouseMoveEventType:
		me := e.(*input.MouseMoveEvent)
		i.mouseX, i.mouseY = me.X, me.Y
	}
}

func (i *IntentMapper) HandlePacket(p *core.InPacket) {
	var id int
	p.Read(&id)
	switch id {
	case KeyEvent:
		var key input.Key
		var scancode int
		var act input.Action
		var mods input.ModifierKey
		p.Read(&key, &scancode, &act, &mods)

		if act == input.Press {
			i.keyStates[int(key)] = input.Pressed

			in := action{input.StickyKeys, int(key), input.Pressed}
			if _, ok := i.inputMap[in]; ok {
				*i.intents[i.inputMap[in]] = true
			}
		} else if act == input.Release {
			i.keyStates[int(key)] = input.Released

			in := action{input.StickyKeys, int(key), input.Released}
			if _, ok := i.inputMap[in]; ok {
				*i.intents[i.inputMap[in]] = true
			}
		}
	case MouseBtnEvent:
		var btn input.MouseButton
		var act input.Action
		var mods input.ModifierKey
		p.Read(&btn, &act, &mods)

		if act == input.Press {
			i.mouseStates[int(btn)] = input.Pressed

			in := action{input.StickyMouseButtons, int(btn), input.Pressed}
			if _, ok := i.inputMap[in]; ok {
				*i.intents[i.inputMap[in]] = true
			}
		} else if act == input.Release {
			i.mouseStates[int(btn)] = input.Released

			in := action{input.StickyMouseButtons, int(btn), input.Released}
			if _, ok := i.inputMap[in]; ok {
				*i.intents[i.inputMap[in]] = true
			}
		}
	case MouseMoveEvent:
		var x, y int
		p.Read(&x, &y)
		i.mouseX, i.mouseY = x, y
	}
}
