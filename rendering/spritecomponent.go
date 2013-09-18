package rend

import (
	"fission/core"
	//glfw "github.com/go-gl/glfw3"
)

// The type bits for SpriteComponent
var SpriteComponentType = core.NextComponentType()

type SpriteComponent struct {
}

func NewSpriteComponent(fileName string) *SpriteComponent {
	return &SpriteComponent{}
}

func (s *SpriteComponent) Serialize() {
}

func (s *SpriteComponent) Deserialize() {
}

func (s *SpriteComponent) TypeBits() core.TypeBits {
	return SpriteComponentType
}

func (s *SpriteComponent) Render(pos *core.Vector2, rot, scale float32) {
}
