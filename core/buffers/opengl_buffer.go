package buffers

import (
	"encoding/binary"
	"github.com/go-gl/gl/v4.1-core/gl"
)

func CreateNewVao(vbo uint32, ebo uint32, attribSize int32) uint32 {
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.VertexAttribPointer(0, attribSize, gl.FLOAT, false, attribSize*int32(binary.Size(float32(0))), nil)

	gl.BindVertexArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)

	return vao
}

func DeleteVao(vaoId uint32) {
	gl.DeleteVertexArrays(1, &vaoId)
}

func CreateNewVbo(vertices []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, binary.Size(float32(0))*len(vertices), gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	return vbo
}

func DeleteVbo(vboId uint32) {
	gl.DeleteBuffers(1, &vboId)
}

func CreateNewEbo(indices []uint32) uint32 {
	var ebo uint32
	gl.GenBuffers(1, &ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, binary.Size(uint32(0))*len(indices), gl.Ptr(indices), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)

	return ebo
}

func DeleteEbo(eboId uint32) {
	gl.DeleteBuffers(1, &eboId)
}
