package core

import (
	"github.com/tedsta/fission/core/event"
)

type System interface {
	Begin(dt float32)
	ProcessEntity(e *Entity, dt float32)
	End(dt float32)
	TypeBits() TypeBits
}

type ActiveEntities struct {
	Entities []*Entity
	typeBits TypeBits
}

func (a *ActiveEntities) Listen(ch chan event.Event) {
	for e := range ch {
		entity := e.(*entityEvent).Ent
		if entity.typeBits&a.typeBits == a.typeBits {
			a.Entities = append(a.Entities, entity)
		}
	}
}
