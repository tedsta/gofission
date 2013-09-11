package fission

type System interface {
	Begin(dt float32)
	ProcessEntity(e *Entity, dt float32)
	End(dt float32)
	TypeBits() int
}

type ActiveEntities struct {
	Entities []*Entity
	typeBits int
}

func (this *ActiveEntities) HandleEvent(event Event) {
	entity := event.(*entityEvent).Ent
	if entity.typeBits&this.typeBits == this.typeBits {
		this.Entities = append(this.Entities, entity)
	}
}