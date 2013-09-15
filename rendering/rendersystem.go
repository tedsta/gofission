package rend

import (
	"fission"
	"github.com/tedsta/go-sfml"
)

const Ptu = 32.0

type RenderSystem struct {
	Window   sfml.RenderWindow
	CamPos   sfml.Vector2f
	CamRot   float32
	CamZoom  float32
	view     sfml.View
	typeBits fission.TypeBits
}

func NewRenderSystem(winTitle string, typeBits fission.TypeBits) *RenderSystem {
	vm := sfml.NewVideoMode(800, 600, 32)
	w := sfml.NewRenderWindowDefault(vm, winTitle)
	return &RenderSystem{w, sfml.NewVector2f(0, 0), 0, 1, w.DefaultView(), typeBits}
}

func (r *RenderSystem) Begin(dt float32) {
	r.Window.Clear(sfml.FromRGB(0, 0, 0))

	r.view.SetCenter(r.CamPos.X()*Ptu, -r.CamPos.Y()*Ptu)
	r.view.SetRotation(-r.CamRot)
	r.view.Zoom(r.CamZoom)
	r.Window.SetView(r.view)
}

func (r *RenderSystem) ProcessEntity(e *fission.Entity, dt float32) {
	transform := e.Component(fission.TransformComponentType).(*fission.TransformComponent)
	pos := sfml.NewVector2f(transform.Pos.X()*Ptu, -transform.Pos.Y()*Ptu)

	renderCmpnts := e.Components(SpriteComponentType | r.typeBits)
	for _, cmpnt := range renderCmpnts {
		cmpnt.(RenderComponent).Render(r.Window, pos, transform.Rot, transform.Scale)
	}
}

func (r *RenderSystem) End(dt float32) {
	r.Window.Display()

	r.view = r.Window.DefaultView()
	r.Window.SetView(r.view)
}

func (r *RenderSystem) TypeBits() fission.TypeBits {
	return fission.TransformComponentType | SpriteComponentType | r.typeBits
}
