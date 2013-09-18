package core

type Vector2 struct {
	X float32
	Y float32
}

func (v *Vector2) Add(other *Vector2) {
	v.X += other.X
	v.Y += other.Y
}

func (v *Vector2) Sub(other *Vector2) {
	v.X -= other.X
	v.Y -= other.Y
}

func (v *Vector2) Mult(s float32) {
	v.X *= s
	v.Y *= s
}

func (v *Vector2) Div(s float32) {
	v.X /= s
	v.Y /= s
}
