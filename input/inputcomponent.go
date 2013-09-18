package input

import (
	"fission/core"
	glfw "github.com/go-gl/glfw3"
)

type InputComponent interface {
	core.Component
	OnKeyPressed(key glfw.Key)
	OnKeyReleased(key glfw.Key)
}
