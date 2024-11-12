package game

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Brick struct {
	Destructable  bool
	HP            int
	PosX, PosY    float64
	Width, Height float64
}

func (b *Brick) Draw(renderer *sdl.Renderer) {
	// Set color based on HP
	switch b.HP {
	case -1:
		renderer.SetDrawColor(120, 120, 120, 255)
	case 0:
		// Invisible
		renderer.SetDrawColor(24, 24, 24, 255)
	case 1:
		renderer.SetDrawColor(255, 0, 0, 255)
	case 2:
		renderer.SetDrawColor(255, 255, 0, 255)
	case 3:
		renderer.SetDrawColor(255, 255, 255, 255)
	}

	// Filled interior
	renderer.FillRectF(&sdl.FRect{
		X: float32(b.PosX),
		Y: float32(b.PosY),
		W: float32(b.Width),
		H: float32(b.Height),
	})

	// Outline
	renderer.SetDrawColor(24, 24, 24, 255)
	renderer.DrawRectF(&sdl.FRect{
		X: float32(b.PosX),
		Y: float32(b.PosY),
		W: float32(b.Width),
		H: float32(b.Height),
	})
}

func (g *Game) InitBricks() {
	brickWidth := float64(g.WindowWidth / 11)
	brickHeight := float64(g.WindowHeight / 24)
	g.Bricks = make([]Brick, 0)
	g.brickCount = 0

	switch g.Stage {
	case 0:
		w := float64(g.WindowWidth) / 4
		h := float64(g.WindowHeight) / 8

		g.Bricks = append(g.Bricks, Brick{
			Destructable: true,
			HP:           1,
			PosX:         float64(g.WindowWidth/2) - (w / 2),
			PosY:         float64(g.WindowHeight/2) - (h / 2),
			Width:        w,
			Height:       h,
		})

		g.brickCount = 1
	case 1: // First stage
		// Single layer of bricks
		Y := float64(g.WindowHeight / 3)
		for l := 0; l < 2; l++ {
			for i := 0; i < 11; i++ {
				g.Bricks = append(g.Bricks, Brick{
					Destructable: true,
					HP:           1,
					PosX:         brickWidth * float64(i),
					PosY:         Y - (brickHeight * float64(l)),
					Width:        brickWidth,
					Height:       brickHeight,
				})

				g.brickCount++
			}
		}
	case 2: // Second stage
		Y := float64(g.WindowHeight / 3)
		for l := 0; l < 3; l++ {
			for i := 0; i < 11; i++ {
				g.Bricks = append(g.Bricks, Brick{
					Destructable: true,
					HP:           1 + 1*l,
					PosX:         brickWidth * float64(i),
					PosY:         Y - (brickHeight * float64(l)),
					Width:        brickWidth,
					Height:       brickHeight,
				})

				g.brickCount++
			}
		}
	}
}
