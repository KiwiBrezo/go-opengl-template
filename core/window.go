package core

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Window struct {
	Window *glfw.Window
}

func NewWindow(width int32, height int32) *Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(int(width), int(height), "Go OpenGL", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	win := Window{Window: window}

	return &win
}
