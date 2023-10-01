package shapes

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"go-open-gl/core/buffers"
	"go-open-gl/models"
)

type Rectangle struct {
	graphicBuffer *buffers.GraphicBuffer
	shaderId      uint32
	SizeX         float32
	SizeY         float32
	Position      models.Vec3
}

func NewRectangle() *Rectangle {
	rectangle := Rectangle{Position: models.Vec3{X: 0.0, Y: 0.0, Z: 0.0}, SizeX: 1.0, SizeY: 1.0}

	rectangle.calculateBuffers()

	return &rectangle
}

func NewRectanglePositioned(position models.Vec3, sizeX float32, sizeY float32) *Rectangle {
	rectangle := Rectangle{Position: position, SizeX: sizeX, SizeY: sizeY}

	rectangle.calculateBuffers()

	return &rectangle
}

func (r *Rectangle) calculateBuffers() {
	vertices := []float32{
		r.Position.X / 2, r.Position.Y / 2, r.Position.Z,
		(r.Position.X + r.SizeX) / 2, r.Position.Y / 2, r.Position.Z,
		(r.Position.X + r.SizeX) / 2, (r.Position.Y + r.SizeY) / 2, r.Position.Z,
		r.Position.X / 2, (r.Position.Y + r.SizeY) / 2, r.Position.Z,
	}

	indices := []uint32{
		0, 3, 2,
		2, 1, 0,
	}

	r.graphicBuffer = buffers.NewGraphicBuffer(vertices, indices, 4, 2)
}

func (r *Rectangle) Bind() {
	gl.BindVertexArray(r.graphicBuffer.VaoId)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, r.graphicBuffer.EboId)
}

func (r *Rectangle) Unbind() {
	gl.BindBuffer(gl.VERTEX_ARRAY_BINDING, 0)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
}

func (r *Rectangle) Draw() {
	r.Bind()

	// Use global shader if not set
	if r.shaderId != 0 {
		gl.UseProgram(r.shaderId)
	}

	gl.DrawElements(gl.TRIANGLES, r.graphicBuffer.IndexSize, gl.UNSIGNED_INT, nil)

	r.Unbind()
}

func (r *Rectangle) Dispose() {
	buffers.DisposeGraphicBuffer(r.graphicBuffer)
}
