package rendering

import (
	"fission"
	"github.com/tedsta/go-sfml"
	"log"
)

// The type bits for SpriteComponent
var SpriteComponentType = fission.NextComponentType()

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

	sizeX, sizeY := texture.Size()
	halfX := float32(sizeX / 2)
	halfY := float32(sizeY / 2)

	spr.SetTexture(texture, false)
	spr.SetOrigin(halfX, halfY)
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

func (this *SpriteComponent) TypeBits() fission.TypeBits {
	return SpriteComponentType
}
