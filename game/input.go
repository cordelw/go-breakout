package game

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type Mouse struct {
	PosX, PosY int32
	State      uint32
}

// Something
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
				g.Ball.VelX = -math.Min(float64(et.X-g.Mouse.PosX)*100, g.Ball.Speed)
				g.Ball.VelY = -g.Ball.Speed
				g.Ball.Held = false
			}
		}

		// Update Mouse trackage
		g.LastMouse = g.Mouse
		g.Mouse.PosX, g.Mouse.PosY, g.Mouse.State = sdl.GetMouseState()
	}
}
