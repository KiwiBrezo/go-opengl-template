package shapes

type Shape2D interface {
	calculateBuffers()
	Bind()
	Unbind()
	Draw()
	Dispose()
}
