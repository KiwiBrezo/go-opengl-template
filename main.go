package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"go-open-gl/core"
	"go-open-gl/models"
	"go-open-gl/models/shapes"
	"log"
	"runtime"
)

var (
	windowWidth  = 800
	windowHeight = 600
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

	mainWindow := core.NewWindow(int32(windowWidth), int32(windowHeight))
	defer glfw.Terminate()

	initOpenGL(int32(windowWidth), int32(windowHeight))

	program, err := core.MakeShaderProgram("assets/shaders/vertex.glsl", "assets/shaders/fragment.glsl")
	if err != nil {
		panic(err)
	}

	shapeList := []shapes.Shape2D{
		shapes.NewCirclePositioned(models.Vec3{X: -0.3, Y: 0.3, Z: 0.0}, 0.1, 64),
		shapes.NewRectanglePositioned(models.Vec3{X: 0.5, Y: -1.0, Z: 0.0}, 0.5, 0.5),
	}

	// For development
	// gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)

	for !mainWindow.Window.ShouldClose() {
		core.Render(mainWindow.Window, program, shapeList)
	}

	log.Println("Disposing programs and buffers")
	gl.DeleteProgram(program)
	for _, shape := range shapeList {
		shape.Dispose()
	}
}
