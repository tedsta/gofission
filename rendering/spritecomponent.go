package rend

import (
	"github.com/tedsta/fission/core"
	"github.com/tedsta/gosfml"
)

// The type bits for SpriteComponent
var SpriteComponentType = core.NextComponentType()

type SpriteComponent struct {
	sprite *sf.Sprite
}

func NewSpriteComponent(fileName string) *SpriteComponent {
	return &SpriteComponent{sf.NewSprite(sf.NewTextureFromFile(fileName))}
}

func (s *SpriteComponent) Serialize() {
}

func (s *SpriteComponent) Deserialize() {
}

func (s *SpriteComponent) TypeBits() core.TypeBits {
	return SpriteComponentType
}

func (s *SpriteComponent) Render(t *sf.RenderTarget, states sf.RenderStates) {
	//states.transform.Combine(transform)
	s.sprite.Render(t, states)
}
