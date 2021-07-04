package game

import "github.com/go-gl/glfw/v3.3/glfw"

func Run() {
	Init()

	window := win.Window

	var dt, lastFrame float32
	lastFrame = float32(glfw.GetTime())
	for !window.ShouldClose() {
		currFrame := float32(glfw.GetTime())
		dt = currFrame - lastFrame
		lastFrame = currFrame
		
		ProcessInput()
		Update(dt)
		Draw()

		glfw.PollEvents()
		window.SwapBuffers()
	}
}
