package rendering

import (
	"fission"
	"github.com/tedsta/go-sfml"
)

type RenderSystem struct {
	Window sfml.RenderWindow
}

func NewRenderSystem(winTitle string) *RenderSystem {
	vm := sfml.NewVideoMode(800, 600, 32)
	return &RenderSystem{sfml.NewRenderWindowDefault(vm, winTitle)}
}

func (this *RenderSystem) Begin(dt float32) {
	this.Window.Drain()
	this.Window.Clear(sfml.FromRGB(0, 0, 0))
}

func (this *RenderSystem) ProcessEntity(e *fission.Entity, dt float32) {
	cmpnts := e.GetComponents(this.TypeBits())

	renderCmpnts := cmpnts[SpriteComponentType]
	for _, cmpnt := range renderCmpnts {
		cmpnt.(RenderComponent).Render(this.Window, sfml.NewVector2f(0, 0))
	}

}

func (this *RenderSystem) End(dt float32) {
	this.Window.Display()
}

func (this *RenderSystem) TypeBits() int {
	return fission.TransformComponentType | SpriteComponentType
}
