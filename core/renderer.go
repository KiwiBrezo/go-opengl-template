package core

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"go-open-gl/models/shapes"
)

func Render(window *glfw.Window, program uint32, shape2DS []shapes.Shape2D) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	for _, shape := range shape2DS {
		shape.Draw()
	}

	glfw.PollEvents()
	window.SwapBuffers()
}
