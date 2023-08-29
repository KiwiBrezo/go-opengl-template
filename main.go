package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"go-3d-tiles-viewer/core"
	"log"
	"runtime"
)

var (
	width  = 800
	height = 600

	vertices = []float32{
		0.5, 0.5, 0.0, // top right
		0.5, -0.5, 0.0, // bottom right
		-0.5, -0.5, 0.0, // bottom left
		-0.5, 0.5, 0.0, // top left
	}

	indices = []uint32{
		0, 1, 3, // first triangle
		1, 2, 3, // second triangle
	}
)

func initOpenGL(width int32, height int32) {
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	gl.Enable(gl.DEPTH_TEST)
	gl.ClearColor(0.2, 0.2, 0.3, 1.0)
	gl.ClearDepth(1)
	gl.DepthFunc(gl.LEQUAL)
	gl.Viewport(0, 0, width, height)
}

func main() {
	runtime.LockOSThread()

	mainWindow := core.NewWindow(int32(width), int32(height))
	defer glfw.Terminate()

	initOpenGL(int32(width), int32(height))

	program, err := core.MakeShaderProgram("assets/shaders/vertex.glsl", "assets/shaders/fragment.glsl")
	if err != nil {
		panic(err)
	}

	vbo := core.MakeVbo(vertices)
	ebo := core.MakeEbo(indices)
	vao := core.MakeVao(vbo, ebo)

	// For development
	//gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)

	for !mainWindow.Window.ShouldClose() {
		core.Render(mainWindow.Window, program, vao)
	}
}
