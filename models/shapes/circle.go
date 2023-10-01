package shapes

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"go-open-gl/core/buffers"
	"go-open-gl/models"
	"math"
)

type Circle struct {
	graphicBuffer *buffers.GraphicBuffer
	shaderId      uint32
	Radius        float32
	SectorCount   uint
	Position      models.Vec3
}

func NewCircle() *Circle {
	circle := Circle{Position: models.Vec3{X: 0.0, Y: 0.0, Z: 0.0}, Radius: 1.0, SectorCount: 8}

	circle.calculateBuffers()

	return &circle
}

func NewCirclePositioned(position models.Vec3, radius float32, sectorCount uint) *Circle {
	circle := Circle{Position: position, Radius: radius, SectorCount: sectorCount}

	circle.calculateBuffers()

	return &circle
}

func (r *Circle) calculateBuffers() {
	const pi = 3.14159265358979323846
	angle := (pi * 2) / float32(r.SectorCount)
	var vertices []float32
	var indices []uint32

	for i := 0; i < int(r.SectorCount); i++ {
		currentAngle := float64(float32(i) * angle)
		vertices = append(vertices, (float32(math.Sin(currentAngle))+r.Position.X)/float32(2), (float32(math.Cos(currentAngle))+r.Position.Y)/float32(2), r.Position.Z)
		indices = append(indices, uint32(r.SectorCount), uint32(i), uint32((i+1)%int(r.SectorCount)))
	}
	vertices = append(vertices, r.Position.Y, r.Position.Y, r.Position.Z)

	r.graphicBuffer = buffers.NewGraphicBuffer(vertices, indices, int32(r.SectorCount), int32(len(indices)/3))
}

func (r *Circle) Bind() {
	gl.BindVertexArray(r.graphicBuffer.VaoId)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, r.graphicBuffer.EboId)
}

func (r *Circle) Unbind() {
	gl.BindBuffer(gl.VERTEX_ARRAY_BINDING, 0)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
}

func (r *Circle) Draw() {
	r.Bind()

	// Use global shader if not set
	if r.shaderId != 0 {
		gl.UseProgram(r.shaderId)
	}

	gl.DrawElements(gl.TRIANGLES, r.graphicBuffer.IndexSize, gl.UNSIGNED_INT, nil)

	r.Unbind()
}

func (r *Circle) Dispose() {
	buffers.DisposeGraphicBuffer(r.graphicBuffer)
}
