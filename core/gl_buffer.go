package core

import (
	"encoding/binary"
	"github.com/go-gl/gl/v4.1-core/gl"
)

func MakeVao(vbo uint32, ebo uint32) uint32 {
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, int32(3*binary.Size(float32(0))), nil)

	return vao
}

func MakeVbo(vertices []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, binary.Size(float32(0))*len(vertices), gl.Ptr(vertices), gl.STATIC_DRAW)

	return vbo
}

func MakeEbo(indices []uint32) uint32 {
	var ebo uint32
	gl.GenBuffers(1, &ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, binary.Size(uint32(0))*len(indices), gl.Ptr(indices), gl.STATIC_DRAW)

	return ebo
}
