package net

import (
	"encoding/gob"
	"github.com/tedsta/fission/core"
	"github.com/tedsta/fission/core/event"
	"github.com/tedsta/fission/input"
)

type IntentSystem struct {
	conn    *Connection
	evt     *event.Manager
	events  []event.Event
	packets map[NetId][]*core.InPacket

	hndId int
}

func NewIntentSystem(conn *Connection, evt *event.Manager) *IntentSystem {
	i := &IntentSystem{conn: conn, evt: evt}
	i.packets = make(map[NetId][]*core.InPacket)
	i.hndId = conn.RegisterHandlerAuto(i.HandlePacket)
	if conn.Type() == Client || conn.Type() == None {
		evt.AddHandler(input.KeyEventType, i)
		evt.AddHandler(input.MouseBtnEventType, i)
		evt.AddHandler(input.MouseMoveEventType, i)
	}
	return i
}

func (i *IntentSystem) Begin(dt float32) {
}

func (i *IntentSystem) ProcessEntity(e *core.Entity, dt float32) {
	intent := e.Component(IntentComponentType).(*IntentComponent)

	for k, _ := range intent.intents {
		intent.intents[k] = false
	}

	for k := 0; k < int(input.KeyLast); k++ {
		if intent.keyStates[k] == input.Pressed {
			intent.keyStates[k] = input.Down
		} else if intent.keyStates[k] == input.Released {
			intent.keyStates[k] = input.Up
		}

		in := action{input.StickyKeys, k, intent.keyStates[k]}

		if _, ok := intent.inputMap[in]; ok {
			intent.intents[intent.inputMap[in]] = true
		}
	}

	for m := 0; m < int(input.MouseButtonLast); m++ {
		if intent.mouseStates[m] == input.Pressed {
			intent.mouseStates[m] = input.Down
		} else if intent.mouseStates[m] == input.Released {
			intent.mouseStates[m] = input.Up
		}

		in := action{input.StickyMouseButtons, m, intent.mouseStates[m]}

		if _, ok := intent.inputMap[in]; ok {
			intent.intents[intent.inputMap[in]] = true
		}
	}

	// ###
	// Event handling
	if (i.conn.Type() == Client && i.conn.NetId() == intent.netId) || i.conn.Type() == None {
		for _, e := range i.events {
			switch e.Type() {
			case input.KeyEventType:
				ke := e.(*input.KeyEvent)

				// Send it across the network
				if ke.Action == input.Press || ke.Action == input.Release {
					packet := core.NewOutPacket(nil)
					packet.Write(intent.netId, KeyEvent, ke.Key, ke.Scancode, ke.Action, ke.Mods)
					i.conn.Send(packet, i.hndId, 0, 0, false)
				}

				if ke.Action == input.Press {
					intent.keyStates[int(ke.Key)] = input.Pressed

					in := action{input.StickyKeys, int(ke.Key), input.Pressed}
					if _, ok := intent.inputMap[in]; ok {
						intent.intents[intent.inputMap[in]] = true
					}
				} else if ke.Action == input.Release {
					intent.keyStates[int(ke.Key)] = input.Released

					in := action{input.StickyKeys, int(ke.Key), input.Released}
					if _, ok := intent.inputMap[in]; ok {
						intent.intents[intent.inputMap[in]] = true
					}
				}
			case input.MouseBtnEventType:
				me := e.(*input.MouseBtnEvent)

				// Send it across the network
				if me.Action == input.Press || me.Action == input.Release {
					packet := core.NewOutPacket(nil)
					packet.Write(intent.netId, MouseBtnEvent, me.Btn, me.Action, me.Mods)
					i.conn.Send(packet, i.hndId, 0, 0, false)
				}

				if me.Action == input.Press {
					intent.mouseStates[int(me.Btn)] = input.Pressed

					in := action{input.StickyMouseButtons, int(me.Btn), input.Pressed}
					if _, ok := intent.inputMap[in]; ok {
						intent.intents[intent.inputMap[in]] = true
					}
				} else if me.Action == input.Release {
					intent.mouseStates[int(me.Btn)] = input.Released

					in := action{input.StickyMouseButtons, int(me.Btn), input.Released}
					if _, ok := intent.inputMap[in]; ok {
						intent.intents[intent.inputMap[in]] = true
					}
				}
			case input.MouseMoveEventType:
				me := e.(*input.MouseMoveEvent)
				intent.mouseX, intent.mouseY = me.X, me.Y
			}
		}
	}

	// ###
	// Packet handling
	for _, p := range i.packets[intent.netId] {
		var id int
		p.Read(&id)
		switch id {
		case KeyEvent:
			var key input.Key
			var scancode int
			var act input.Action
			var mods input.ModifierKey
			p.Read(&key, &scancode, &act, &mods)

			// Forward controls to other clients
			if i.conn.Type() == Server && (act == input.Press || act == input.Release) {
				packet := core.NewOutPacket(nil)
				packet.Write(intent.netId, KeyEvent, key, scancode, act, mods)
				i.conn.Send(packet, i.hndId, 0, intent.netId, false)
			}

			if act == input.Press {
				intent.keyStates[int(key)] = input.Pressed

				in := action{input.StickyKeys, int(key), input.Pressed}
				if _, ok := intent.inputMap[in]; ok {
					intent.intents[intent.inputMap[in]] = true
				}
			} else if act == input.Release {
				intent.keyStates[int(key)] = input.Released

				in := action{input.StickyKeys, int(key), input.Released}
				if _, ok := intent.inputMap[in]; ok {
					intent.intents[intent.inputMap[in]] = true
				}
			}
		case MouseBtnEvent:
			var btn input.MouseButton
			var act input.Action
			var mods input.ModifierKey
			p.Read(&btn, &act, &mods)

			// Forward controls to other clients
			if i.conn.Type() == Server && (act == input.Press || act == input.Release) {
				packet := core.NewOutPacket(nil)
				packet.Write(intent.netId, MouseBtnEvent, btn, act, mods)
				i.conn.Send(packet, i.hndId, 0, intent.netId, false)
			}

			if act == input.Press {
				intent.mouseStates[int(btn)] = input.Pressed

				in := action{input.StickyMouseButtons, int(btn), input.Pressed}
				if _, ok := intent.inputMap[in]; ok {
					intent.intents[intent.inputMap[in]] = true
				}
			} else if act == input.Release {
				intent.mouseStates[int(btn)] = input.Released

				in := action{input.StickyMouseButtons, int(btn), input.Released}
				if _, ok := intent.inputMap[in]; ok {
					intent.intents[intent.inputMap[in]] = true
				}
			}
		case MouseMoveEvent:
			var x, y int
			p.Read(&x, &y)
			intent.mouseX, intent.mouseY = x, y
		}
	}
}

func (i *IntentSystem) End(dt float32) {
	i.events = nil
	i.packets = make(map[NetId][]*core.InPacket)
}

func (i *IntentSystem) TypeBits() (core.TypeBits, core.TypeBits) {
	return IntentComponentType, 0
}

func (i *IntentSystem) HandleEvent(e event.Event) {
	i.events = append(i.events, e)
}

func (i *IntentSystem) HandlePacket(p *core.InPacket) {
	var netId NetId
	p.Read(&netId)
	i.packets[netId] = append(i.packets[netId], p)
}

// #############################################################################

func RegisterComponents() {
	IntentComponentType = core.RegisterComponent(IntentComponentFactory)

	gob.Register(action{})
}
