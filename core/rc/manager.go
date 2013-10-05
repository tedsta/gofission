package rc

type Id uint

var nextId Id

func NextId() Id {
	defer func() { nextId++ }()
	return nextId
}

var res []interface{} // Resources

func Add(id Id, rc interface{}) {
	if int(id) >= len(res) { // Check if we have enough room
		// Resize the resource table accordingly
		newRes := make([]interface{}, id+1)
		copy(newRes, res)
		res = newRes
	}

	res[id] = rc
}

func Get(id Id) interface{} {
	if int(id) >= len(res) {
		panic("Trying to access undefined resource.")
	}
	return res[id]
}
