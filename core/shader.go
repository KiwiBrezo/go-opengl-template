package core

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"go-3d-tiles-viewer/io"
	"log"
	"strings"
)

func CompileShader(shaderSource string, shaderType uint32) (shaderId uint32, err error) {
	shaderId = gl.CreateShader(shaderType)

	sources, free := gl.Strs(shaderSource)
	gl.ShaderSource(shaderId, 1, sources, nil)
	free()
	gl.CompileShader(shaderId)

	var status int32
	gl.GetShaderiv(shaderId, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shaderId, gl.INFO_LOG_LENGTH, &logLength)

		errorLog := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shaderId, logLength, nil, gl.Str(errorLog))

		return 0, fmt.Errorf("(core.CompileShader) Failed to compile shader \n %v \n ErrorLog: %v", shaderSource, errorLog)
	}

	return shaderId, nil
}

func MakeShaderProgram(vertexSourcePath string, fragmentSourcePath string) (program uint32, err error) {
	vertexShaderSource, err := io.ReadShaderFile(vertexSourcePath)
	if err != nil {
		log.Printf("(core.MakeShaderProgram) There was a problem reading a vertex shader from path: %s \n %v", vertexShaderSource, err)
		return 0, err
	}

	log.Println("(core.MakeShaderProgram) Vertex shader:")
	fmt.Println(vertexShaderSource)

	vertexShader, err := CompileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		log.Printf("(core.MakeShaderProgram) There was a problem compiling the vertex shader from path: %s \n %v", vertexShaderSource, err)
		return 0, err
	}

	fragmentShaderSource, err := io.ReadShaderFile(fragmentSourcePath)
	if err != nil {
		log.Printf("(core.MakeShaderProgram) There was a problem reading a fragment shader from path: %s \n %v", vertexShaderSource, err)
		return 0, err
	}

	log.Println("(core.MakeShaderProgram) Fragment shader:")
	fmt.Println(fragmentShaderSource)

	fragmentShader, err := CompileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		log.Printf("(core.MakeShaderProgram) There was a problem compiling the fragment shader from path: %s \n %v", vertexShaderSource, err)
		return 0, err
	}

	program = gl.CreateProgram()
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	return program, nil
}
