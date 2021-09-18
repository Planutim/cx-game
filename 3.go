package main

import (
	"runtime"

	"github.com/Planutim/myopengl"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func init() {
	runtime.LockOSThread()
}
func main() {
	window := myopengl.InitGlfw(800, 600)

	gl.ClearColor(0.2, 0.5, 0.1, 1)
	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)

		glfw.PollEvents()
		window.SwapBuffers()
	}
}
