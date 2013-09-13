package fission

var (
	EventAddEntity    = 0
	EventRemoveEntity = 1
)

type Scene struct {
	entities  []*Entity
	nextIndex int // Used for iterating through the scene's entities
}

func (s *Scene) AddEntity(ent *Entity) {
	s.entities = append(s.entities, ent)
}

func (s *Scene) FindEntity(id int) *Entity {
	for _, ent := range s.entities {
		if ent.id == id {
			return ent
		}
	}
	return nil
}

func (s *Scene) RemoveEntity(id int) {
	for i, ent := range s.entities {
		if ent.id == id {
			s.entities[i] = s.entities[len(s.entities)-1]
			s.entities = s.entities[:len(s.entities)-2]
			return
		}
	}
}

func (s *Scene) Save(fileName string) {

}

func (s *Scene) Load(fileName string) {

}

// BeginEntities initializes an iteration through the entities in a scene
func (s *Scene) BeginEnt() *Entity {
	s.nextIndex = 0
	if len(s.entities) > 0 {
		return s.entities[s.nextIndex]
	}
	return nil
}

// NextEntity returns the next entity when iterating through the entities
func (s *Scene) NextEntity() *Entity {
	if int(len(s.entities)) > s.nextIndex+1 {
		s.nextIndex++
		return s.entities[s.nextIndex]
	}
	return nil
}
