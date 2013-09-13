package rendering

import (
	"fission"
	"github.com/tedsta/go-sfml"
)

const Ptu = 32.0

type RenderSystem struct {
	Window  sfml.RenderWindow
	CamPos  sfml.Vector2f
	CamRot  float32
	CamZoom float32
	view    sfml.View
}

func NewRenderSystem(winTitle string) *RenderSystem {
	vm := sfml.NewVideoMode(800, 600, 32)
	w := sfml.NewRenderWindowDefault(vm, winTitle)
	return &RenderSystem{w, sfml.NewVector2f(0, 0), 0, 1, w.DefaultView()}
}

func (this *RenderSystem) Begin(dt float32) {
	this.Window.Drain()
	this.Window.Clear(sfml.FromRGB(0, 0, 0))

	this.view.SetCenter(this.CamPos.X()*Ptu, -this.CamPos.Y()*Ptu)
	this.view.SetRotation(-this.CamRot)
	this.view.Zoom(this.CamZoom)
	this.Window.SetView(this.view)
}

func (this *RenderSystem) ProcessEntity(e *fission.Entity, dt float32) {
	transform := e.Component(fission.TransformComponentType).(*fission.TransformComponent)
	pos := sfml.NewVector2f(transform.Pos.X()*Ptu, -transform.Pos.Y()*Ptu)

	renderCmpnts := e.Components(SpriteComponentType)
	for _, cmpnt := range renderCmpnts {
		cmpnt.(RenderComponent).Render(this.Window, pos)
	}
}

func (this *RenderSystem) End(dt float32) {
	this.Window.Display()

	this.view = this.Window.DefaultView()
	this.Window.SetView(this.view)
}

func (this *RenderSystem) TypeBits() fission.TypeBits {
	return fission.TransformComponentType | SpriteComponentType
}
