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

func (g *Game) HandleInput() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch et := event.(type) {

		// Close button
		case *sdl.QuitEvent:
			g.Active = false

		// Move the player's paddle based
		// on mouse position
		case *sdl.MouseMotionEvent:
			g.Paddle.PosX = float64(et.X) - float64(g.Paddle.Width/2)

		// Release ball if lmb clicked while currently held
		case *sdl.MouseButtonEvent:
			if et.Button == sdl.BUTTON_LEFT && g.Ball.Held {
				g.Ball.VelY = -g.Ball.Speed
				g.Ball.Held = false

				dir := float64(et.X - g.Mouse.PosX)
				if dir != 0 {
					g.Ball.VelX = -math.Min(dir*10, g.Ball.Speed)
				} else {
					if rand.Float32() < 0.5 {
						g.Ball.VelX = -g.Ball.Speed
					} else {
						g.Ball.VelX = g.Ball.Speed
					}
				}
			}

		case *sdl.KeyboardEvent:
			if et.Keysym.Scancode == sdl.SCANCODE_SPACE {
				g.Ball.Held = true
			}
		}

		// Update Mouse trackage
		g.LastMouse = g.Mouse
		g.Mouse.PosX, g.Mouse.PosY, g.Mouse.State = sdl.GetMouseState()
	}
}
