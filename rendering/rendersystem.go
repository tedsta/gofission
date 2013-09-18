package rend

import (
	"fission/core"
	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
)

const Ptu = 32.0

type RenderSystem struct {
	Window   *glfw.Window
	CamPos   *core.Vector2
	CamRot   float32
	CamScale float32
	typeBits core.TypeBits
}

func NewRenderSystem(winTitle string, typeBits core.TypeBits) *RenderSystem {
	w, err := glfw.CreateWindow(800, 600, winTitle, nil, nil)
	if err != nil {
		panic(err)
	}

	w.MakeContextCurrent()
	return &RenderSystem{w, &core.Vector2{}, 0, 1, typeBits}
}

func (r *RenderSystem) Begin(dt float32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (r *RenderSystem) ProcessEntity(e *core.Entity, dt float32) {
	trans := e.Component(core.TransformComponentType).(*core.TransformComponent)
	pos := &core.Vector2{trans.Pos.X, -trans.Pos.Y}
	pos.Mult(Ptu)

	rot := trans.Rot - r.CamRot
	scale := trans.Scale * r.CamScale

	renderCmpnts := e.Components(SpriteComponentType | r.typeBits)
	for _, cmpnt := range renderCmpnts {
		cmpnt.(RenderComponent).Render(pos, rot, scale)
	}
}

func (r *RenderSystem) End(dt float32) {
	r.Window.SwapBuffers()
}

func (r *RenderSystem) TypeBits() core.TypeBits {
	return core.TransformComponentType | SpriteComponentType | r.typeBits
}
