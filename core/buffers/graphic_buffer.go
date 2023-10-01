package buffers

type GraphicBuffer struct {
	VaoId       uint32
	VboId       uint32
	EboId       uint32
	VertexCount int32
	IndexCount  int32
	VertexSize  int32
	IndexSize   int32
}

func NewGraphicBuffer(vertices []float32, indices []uint32, vertexCount int32, indexCount int32) *GraphicBuffer {
	vertexSize := len(vertices)
	indexSize := len(indices)
	vboId := CreateNewVbo(vertices)
	eboId := CreateNewEbo(indices)
	vaoId := CreateNewVao(vboId, eboId, int32(indexSize)/indexCount)

	gBuff := GraphicBuffer{
		VaoId:       vaoId,
		VboId:       vboId,
		EboId:       eboId,
		VertexCount: vertexCount,
		IndexCount:  indexCount,
		VertexSize:  int32(vertexSize),
		IndexSize:   int32(indexSize),
	}

	return &gBuff
}

func DisposeGraphicBuffer(gBuff *GraphicBuffer) {
	DeleteVao(gBuff.VaoId)
	DeleteVbo(gBuff.VboId)
	DeleteEbo(gBuff.EboId)
}
