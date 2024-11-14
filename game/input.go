package game

import (
	"math"
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

type Mouse struct {
	PosX, PosY int32
	State      uint32
}

func (g *Game) handleInput() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch et := event.(type) {

		// Close button
		case *sdl.QuitEvent:
			g.Active = false

		// Move the player's paddle based
		// on mouse position
		case *sdl.MouseMotionEvent:
			g.paddle.PosX = float64(et.X) - float64(g.paddle.Width/2)

		// Release ball if lmb clicked while currently held
		case *sdl.MouseButtonEvent:
			if et.Button == sdl.BUTTON_LEFT && g.ball.Held {
				g.ball.VelY = -g.ball.Speed
				g.ball.Held = false

				dir := float64(et.X - g.mouse.PosX)
				if dir != 0 {
					g.ball.VelX = -math.Min(dir*10, g.ball.Speed/2)
				} else {
					if rand.Float32() < 0.5 {
						g.ball.VelX = -g.ball.Speed / 2
					} else {
						g.ball.VelX = g.ball.Speed / 2
					}
				}
			}

			/*
				case *sdl.KeyboardEvent:
					if et.Keysym.Scancode == sdl.SCANCODE_SPACE {
						g.ball.Held = true
					}
			*/
		}

		// Update Mouse trackage
		g.lastMouse = g.mouse
		g.mouse.PosX, g.mouse.PosY, g.mouse.State = sdl.GetMouseState()
	}
}
