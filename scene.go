package fission

var (
	EventAddEntity    = 0
	EventRemoveEntity = 1
)

type Scene struct {
	entities  []*Entity
	nextIndex int // Used for iterating through the scene's entities
}

func (this *Scene) AddEntity(ent *Entity) {
	this.entities = append(this.entities, ent)
}

func (this *Scene) FindEntity(id int) *Entity {
	for _, ent := range this.entities {
		if ent.id == id {
			return ent
		}
	}
	return nil
}

func (this *Scene) RemoveEntity(id int) {
	for i, ent := range this.entities {
		if ent.id == id {
			this.entities[i] = this.entities[len(this.entities)-1]
			this.entities = this.entities[:len(this.entities)-2]
			return
		}
	}
}

func (this *Scene) Save(fileName string) {

}

func (this *Scene) Load(fileName string) {

}

// BeginEntities initializes an iteration through the entities in a scene
func (this *Scene) BeginEnt() *Entity {
	this.nextIndex = 0
	if len(this.entities) > 0 {
		return this.entities[this.nextIndex]
	}
	return nil
}

// NextEntity returns the next entity when iterating through the entities
func (this *Scene) NextEntity() *Entity {
	if int(len(this.entities)) > this.nextIndex+1 {
		this.nextIndex++
		return this.entities[this.nextIndex]
	}
	return nil
}
