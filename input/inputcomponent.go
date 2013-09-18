package input

import (
	glfw "github.com/go-gl/glfw3"
	"github.com/tedsta/fission/core"
)

type InputComponent interface {
	core.Component
	OnKeyPressed(key glfw.Key)
	OnKeyReleased(key glfw.Key)
}
