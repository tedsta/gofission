package rnd

import (
	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	"github.com/tedsta/fission/core"
	"github.com/tedsta/gosfml"
)

const Ptu = 32.0

type RenderSystem struct {
	Window *glfw.Window
	Target *sf.RenderTarget
}

func NewRenderSystem(sizeX, sizeY int, winTitle string) *RenderSystem {
	if !glfw.Init() {
		panic("Can't init glfw!")
	}
	gl.Init()

	w, err := glfw.CreateWindow(sizeX, sizeY, winTitle, nil, nil)
	if err != nil {
		panic(err)
	}
	w.MakeContextCurrent()

	rt := sf.NewRenderTarget(sf.Vector2{float32(sizeX), float32(sizeY)})
	r := &RenderSystem{w, rt}

	w.SetFramebufferSizeCallback(r.onResize)

	return r
}

func (r *RenderSystem) Begin(dt float32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (r *RenderSystem) ProcessEntity(e *core.Entity, dt float32) {
	trans := e.Component(TransformComponentType).(*TransformComponent)

	rs := sf.RenderStates{sf.BlendAlpha, trans.T.Transform(), nil}

	renderCmpnts := e.Components(RenderComponentType)
	for _, cmpnt := range renderCmpnts {
		render := cmpnt.(*RenderComponent).Render
		if render != nil {
			render(r.Target, rs)
		}
	}
}

func (r *RenderSystem) End(dt float32) {
	r.Window.SwapBuffers()
}

func (r *RenderSystem) TypeBits() core.TypeBits {
	return TransformComponentType | RenderComponentType
}

// Callbacks ###################################################################

func (r *RenderSystem) onResize(wnd *glfw.Window, w, h int) {
	r.Target.Size.X = float32(w)
	r.Target.Size.Y = float32(h)
}
