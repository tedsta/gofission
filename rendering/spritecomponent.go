package rendering

import (
	"github.com/tedsta/go-sfml"
	"log"
)

const SpriteComponentType = 2

type SpriteComponent struct {
	sprite sfml.Sprite
}

func NewSpriteComponent(fileName string) *SpriteComponent {
	spr, err := sfml.NewSprite()
	if err != nil {
		log.Fatal(err)
	}

	texture, err := sfml.TextureFromFile(fileName, sfml.IntRect{nil})
	if err != nil {
		log.Fatal(err)
	}

	spr.SetTexture(texture, false)
	return &SpriteComponent{spr}
}

func (this *SpriteComponent) Serialize() {
}

func (this *SpriteComponent) Deserialize() {

}

func (this *SpriteComponent) Render(window sfml.RenderWindow, pos sfml.Vector2f) {
	this.sprite.SetPosition(pos.X(), pos.Y())
	window.DrawSpriteDefault(this.sprite)
}

func (this *SpriteComponent) TypeBits() int {
	return SpriteComponentType
}
