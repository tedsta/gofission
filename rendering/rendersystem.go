package rnd

import (
	glfw "github.com/go-gl/glfw3"
	"github.com/tedsta/fission/core"
	"github.com/tedsta/gosfml"
)

const Ptu = 32.0

type RenderSystem struct {
	Window  *glfw.Window
	Target  *sf.RenderTarget
	View    *sf.View
	BgColor sf.Color // Background color
	optBits core.TypeBits
}

func NewRenderSystem(sizeX, sizeY int, winTitle string, optBits core.TypeBits) *RenderSystem {
	w, err := glfw.CreateWindow(sizeX, sizeY, winTitle, nil, nil)
	if err != nil {
		panic(err)
	}
	w.MakeContextCurrent()

	rt := sf.NewRenderTarget(sf.Vector2{float32(sizeX), float32(sizeY)})
	view := rt.DefaultView()
	r := &RenderSystem{w, rt, &view, sf.Color{0, 0, 0, 0}, SpriteComponentType | optBits}

	w.SetFramebufferSizeCallback(r.onResize)

	return r
}

func (r *RenderSystem) Begin(dt float32) {
	r.Target.Clear(r.BgColor)
	r.Target.SetView(*r.View)
}

func (r *RenderSystem) ProcessEntity(e *core.Entity, dt float32) {
	trans := e.Component(TransformComponentType).(*TransformComponent)

	rs := sf.RenderStates{sf.BlendAlpha, trans.T.Transform(), nil}

	renderCmpnts := e.Components(r.optBits)
	for _, cmpnt := range renderCmpnts {
		render := cmpnt.(RenderComponent).Render
		if render != nil {
			render(r.Target, rs)
		}
	}
}

func (r *RenderSystem) End(dt float32) {
	r.Window.SwapBuffers()
}

func (r *RenderSystem) TypeBits() (core.TypeBits, core.TypeBits) {
	return TransformComponentType, r.optBits
}

// init ##########################################################

func init() {
	SpriteComponentType = core.RegisterComponent(SpriteComponentFactory)
	TransformComponentType = core.RegisterComponent(TransformComponentFactory)
}

// Callbacks ###################################################################

func (r *RenderSystem) onResize(wnd *glfw.Window, w, h int) {
	r.Target.Size.X = float32(w)
	r.Target.Size.Y = float32(h)
}
