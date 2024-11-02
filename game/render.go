package game

import "github.com/veandco/go-sdl2/sdl"

func (g *Game) Draw() {
	// Clear draw buffer
	g.Renderer.SetDrawColor(24, 24, 24, 255)
	g.Renderer.Clear()

	// Draw paddle and ball(s)
	g.Paddle.Draw(g.Renderer)
	g.Ball.Draw(g.Renderer)

	// Present renderer
	g.Renderer.Present()
}

func DrawCircle(renderer *sdl.Renderer, posX, posY, radius int) {
	diameter := radius * 2

	var x, y int = radius - 1, 0
	var tx, ty int = 1, 1
	error := tx - diameter

	for x >= y {
		xpx, xmx := posX+x, posX-x
		xpy, xmy := posX+y, posX-y
		ypy, ymy := posY+y, posY-y
		ypx, ymx := posY+x, posY-x

		renderer.DrawPoint(int32(xpx), int32(ymy))
		renderer.DrawPoint(int32(xpx), int32(ypy))
		renderer.DrawPoint(int32(xmx), int32(ymy))
		renderer.DrawPoint(int32(xmx), int32(ypy))
		renderer.DrawPoint(int32(xpy), int32(ymx))
		renderer.DrawPoint(int32(xpy), int32(ypx))
		renderer.DrawPoint(int32(xmy), int32(ymx))
		renderer.DrawPoint(int32(xmy), int32(ypx))

		if error <= 0 {
			y++
			error += ty
			ty += 2
		}
		if error > 0 {
			x--
			tx += 2
			error += tx - diameter
		}
	}
}
