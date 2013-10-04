package rc

type Id uint

var nextId Id

func NextId() Id {
	defer func() { nextId++ }()
	return nextId
}

type Manager struct {
	res []interface{} // Resources
}

func (m *Manager) Add(id Id, rc interface{}) {
	m.res[id] = rc
}

func (m *Manager) Get(id Id) interface{} {
	return m.res[id]
}
