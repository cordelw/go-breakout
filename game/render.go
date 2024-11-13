package game

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func (g *Game) Draw() {
	// Clear draw buffer
	g.Renderer.SetDrawColor(24, 24, 24, 255)
	g.Renderer.Clear()

	// Draw Bricks
	for _, b := range g.Bricks {
		b.Draw(g.Renderer)
	}

	// Draw paddle and ball(s)
	g.Paddle.Draw(g.Renderer)
	g.Ball.Draw(g.Renderer)

	// Start game screen
	if g.Stage == 0 {
		// Breakout text
		tw := (g.WindowWidth / 3) * 2
		th := (g.WindowHeight) / 8
		dst := &sdl.Rect{
			X: (g.WindowWidth / 2) - (tw / 2),
			Y: (g.WindowHeight / 4) - (th / 2),
			W: tw,
			H: th,
		}
		g.Renderer.Copy(
			g.textures["breakout"],
			nil,
			dst,
		)

		// Start Brick
		cbrick := g.Bricks[0]
		g.Renderer.Copy(
			g.textures["start"],
			nil,
			&sdl.Rect{
				X: int32(cbrick.PosX),
				Y: int32(cbrick.PosY),
				W: int32(cbrick.Width),
				H: int32(cbrick.Height),
			},
		)
	}

	// Game over screen
	if g.Stage == 999 {
		// Game over text in top quarter of screen
		tw := (g.WindowWidth / 3) * 2
		th := (g.WindowHeight) / 8
		dst := &sdl.Rect{
			X: (g.WindowWidth / 2) - (tw / 2),
			Y: (g.WindowHeight / 4) - (th / 2),
			W: tw,
			H: th,
		}
		g.Renderer.Copy(
			g.textures["game over"],
			nil,
			dst,
		)

		// Restart on brick
		cbrick := g.Bricks[0]
		g.Renderer.Copy(
			g.textures["restart"],
			nil,
			&sdl.Rect{
				X: int32(cbrick.PosX),
				Y: int32(cbrick.PosY),
				W: int32(cbrick.Width),
				H: int32(cbrick.Height),
			},
		)
	}

	// Win screen
	if g.Stage == 6 {
		tw := (g.WindowWidth / 3) * 2
		th := (g.WindowHeight) / 8
		dst := &sdl.Rect{
			X: (g.WindowWidth / 2) - (tw / 2),
			Y: (g.WindowHeight / 4) - (th / 2),
			W: tw,
			H: th,
		}
		g.Renderer.Copy(
			g.textures["congratulations"],
			nil,
			dst,
		)
	}

	// Points and ballcount display
	if g.Stage != 0 {
		tw := g.WindowHeight / 20
		th := g.WindowHeight / 15
		var textSurface *sdl.Surface

		// Draw points
		//
		pcdst := &sdl.Rect{
			X: 0,
			Y: 0,
			W: tw * 6,
			H: th,
		}
		g.Renderer.Copy(g.textures["score"], nil, pcdst)

		// count
		pcstr := fmt.Sprint(g.points)
		textSurface, _ = g.font.RenderUTF8Solid(
			pcstr,
			sdl.Color{
				R: 255,
				G: 255,
				B: 255,
			},
		)
		pctext, _ := g.Renderer.CreateTextureFromSurface(textSurface)
		textSurface.Free()

		g.Renderer.Copy(pctext, nil, &sdl.Rect{
			X: pcdst.W + tw,
			Y: 0,
			W: int32(len(pcstr)) * tw,
			H: th,
		})
		pctext.Destroy()

		// Draw balls
		//
		if g.Stage != 999 && g.Stage != 6 {
			bcdst := &sdl.Rect{
				X: 0,
				Y: th,
				W: tw * 6,
				H: th,
			}
			g.Renderer.Copy(g.textures["balls"], nil, bcdst)

			// count
			bcstr := fmt.Sprint(g.ballCount)
			textSurface, _ = g.font.RenderUTF8Solid(
				bcstr,
				sdl.Color{
					R: 255,
					G: 255,
					B: 255,
				},
			)
			bctext, _ := g.Renderer.CreateTextureFromSurface(textSurface)
			textSurface.Free()

			g.Renderer.Copy(bctext, nil, &sdl.Rect{
				X: bcdst.W + tw,
				Y: th,
				W: int32(len(bcstr)) * tw,
				H: th,
			})
			bctext.Destroy()

		}
	}

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
