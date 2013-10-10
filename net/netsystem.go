package net

import (
	"enet"
	"github.com/tedsta/fission/core"
)

type NetSystem struct {
	conn *Connection
}

func NewNetSystem() *NetSystem {
	enet.Init()
	return &NetSystem{}
}

func (n *NetSystem) Begin(dt float32) {
}

func (n *NetSystem) ProcessEntity(e *core.Entity, dt float32) {
}

func (n *NetSystem) End(dt float32) {
}

func (n *NetSystem) TypeBits() (core.TypeBits, core.TypeBits) {
	return 0, 0
}

// #############################################################################

func RegisterComponents() {
	IntentComponentType = core.RegisterComponent(IntentComponentFactory)
}
