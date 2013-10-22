package rc

var res map[string]interface{} // Resources

func Add(name string, rc interface{}) {
	res[name] = rc
}

func Get(name string) interface{} {
	if rc, ok := res[name]; ok {
		return rc
	}
	panic("Attempt to access undefined resource.")
}

func init() {
	res = make(map[string]interface{})
}
