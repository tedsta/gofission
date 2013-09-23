package rnd

import (
	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	"github.com/tedsta/fission/core"
	"github.com/tedsta/gosfml"
)

const Ptu = 32.0

type RenderSystem struct {
	Window   *glfw.Window
	Target   *sf.RenderTarget
	typeBits core.TypeBits
}

func NewRenderSystem(winTitle string, typeBits core.TypeBits) *RenderSystem {
	if !glfw.Init() {
		panic("Can't init glfw!")
	}
	gl.Init()

	w, err := glfw.CreateWindow(800, 600, winTitle, nil, nil)
	if err != nil {
		panic(err)
	}
	w.MakeContextCurrent()

	rt := sf.NewRenderTarget()

	return &RenderSystem{w, rt, typeBits}
}

func (r *RenderSystem) Begin(dt float32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (r *RenderSystem) ProcessEntity(e *core.Entity, dt float32) {
	trans := e.Component(TransformComponentType).(*TransformComponent)

	rs := sf.RenderStates{sf.BlendAlpha, trans.T.Transform(), nil}

	renderCmpnts := e.Components(SpriteComponentType | r.typeBits)
	for _, cmpnt := range renderCmpnts {
		cmpnt.(RenderComponent).Render(r.Target, rs)
	}
}

func (r *RenderSystem) End(dt float32) {
	r.Window.SwapBuffers()
}

func (r *RenderSystem) TypeBits() core.TypeBits {
	return TransformComponentType | SpriteComponentType | r.typeBits
}
