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

	switch g.Stage {
	case 0: // First stage
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
			}
		}
	case 1: // Second stage
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
			}
		}
	}
}

func (g *Game) removeBrick(index int) {
	g.Bricks[index] = g.Bricks[len(g.Bricks)-1]
	g.Bricks = g.Bricks[:len(g.Bricks)-1]
}

func (g *Game) updateBricks() {
	ball := &g.Ball

	// For every brick on field
	for i, b := range g.Bricks {
		// Don't update dead bricks
		if b.HP == 0 {
			continue
		}

		/* COLLISION CHECKS */
		// Inside Horizontal column
		if ball.PosX+float64(ball.Radius) > b.PosX && ball.PosX-float64(ball.Radius) < b.PosX+b.Width {
			// Top
			if ball.PosY+float64(ball.Radius) > b.PosY && ball.PosY < b.PosY+(b.Height/2) {
				ball.PosY = b.PosY - float64(ball.Radius)
				ball.VelY = -ball.VelY
				g.Bricks[i].HP -= 1
			}

			// Bottom
			if ball.PosY-float64(ball.Radius) < b.PosY+b.Height && ball.PosY > b.PosY+(b.Height/2) {
				ball.PosY = b.PosY + b.Height + float64(ball.Radius)
				ball.VelY = -ball.VelY
				g.Bricks[i].HP -= 1
			}
		}

		// Inside vertical column
		if ball.PosY+float64(ball.Radius) > b.PosY && ball.PosY-float64(ball.Radius) < b.PosY+b.Height {
			// Left
			if ball.PosX+float64(ball.Radius) > b.PosX && ball.PosX < b.PosX+(b.Width/2) {
				ball.PosX = b.PosX - float64(ball.Radius)
				ball.VelX = -ball.VelX
				g.Bricks[i].HP -= 1
			}

			// Right
			if ball.PosX-float64(ball.Radius) < b.PosX+b.Width && ball.PosX > b.PosX+(b.Width/2) {
				ball.PosX = b.PosX + b.Width + float64(ball.Radius)
				ball.VelX = -ball.VelX
				g.Bricks[i].HP -= 1
			}
		}
	}
}
