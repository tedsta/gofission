package rend

import (
	"github.com/tedsta/fission/core"
)

// The type bits for SpriteComponent
var SpriteComponentType = core.NextComponentType()

type SpriteComponent struct {
	sprite *Sprite
}

func NewSpriteComponent(fileName string) *SpriteComponent {
	return &SpriteComponent{NewSprite(NewTextureFromFile(fileName))}
}

func (s *SpriteComponent) Serialize() {
}

func (s *SpriteComponent) Deserialize() {
}

func (s *SpriteComponent) TypeBits() core.TypeBits {
	return SpriteComponentType
}

func (s *SpriteComponent) Render(t *RenderTarget, states RenderStates, pos Vector2, rot, scale float32) {
	s.sprite.Render(t, states)
}
