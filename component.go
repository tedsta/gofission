package fission

type Component interface {
	Serialize()
	Deserialize()
	TypeBits() int
}
