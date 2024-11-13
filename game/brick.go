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

func (g *Game) initBricks() {
	brickWidth := float64(g.windowWidth / 11)
	brickHeight := float64(g.windowHeight / 24)
	g.bricks = make([]Brick, 0)
	g.brickCount = 0

	ystart := 2 * (float64(g.windowHeight) / 15)

	switch g.stage {
	case 0: // Start Menu
		w := float64(g.windowWidth) / 4
		h := float64(g.windowHeight) / 8

		g.bricks = append(g.bricks, Brick{
			Destructable: true,
			HP:           1,
			PosX:         float64(g.windowWidth/2) - (w / 2),
			PosY:         float64(g.windowHeight/2) - (h / 2),
			Width:        w,
			Height:       h,
		})

		g.brickCount = 1

	case 999: // Game over menu
		w := float64(g.windowWidth) / 4
		h := float64(g.windowHeight) / 8

		g.bricks = append(g.bricks, Brick{
			Destructable: true,
			HP:           1,
			PosX:         float64(g.windowWidth/2) - (w / 2),
			PosY:         float64(g.windowHeight/2) - (h / 2),
			Width:        w,
			Height:       h,
		})

		g.brickCount = 1

	case 1: // First stage
		// Single layer of bricks
		//Y := float64(g.windowHeight / 3)
		for l := 0; l < 2; l++ {
			for i := 0; i < 11; i++ {
				g.bricks = append(g.bricks, Brick{
					Destructable: true,
					HP:           1,
					PosX:         brickWidth * float64(i),
					PosY:         ystart + brickHeight*float64(l),
					Width:        brickWidth,
					Height:       brickHeight,
				})

				g.brickCount++
			}
		}
	case 2: // Second stage
		for l := 0; l < 3; l++ {
			for i := 0; i < 11; i++ {
				g.bricks = append(g.bricks, Brick{
					Destructable: true,
					HP:           2 + 1 - l,
					PosX:         brickWidth * float64(i),
					PosY:         ystart + (brickHeight * float64(l)),
					Width:        brickWidth,
					Height:       brickHeight,
				})

				g.brickCount++
			}
		}

	case 3: // Stage 3
		for l := 0; l < 2; l++ { // top 2 layers
			for i := 0; i < 11; i++ {
				g.bricks = append(g.bricks, Brick{
					Destructable: true,
					HP:           1 + 1 - l,
					PosX:         brickWidth * float64(i),
					PosY:         ystart + (brickHeight*2)*float64(l),
					Width:        brickWidth,
					Height:       brickHeight,
				})

				g.brickCount++
			}
		}

		// Unbreakable blocks

		g.bricks = append(g.bricks, Brick{
			Destructable: false,
			HP:           -1,
			PosX:         brickWidth * 2,
			PosY:         ystart + brickHeight*4,
			Width:        brickWidth,
			Height:       brickHeight,
		})
		g.bricks = append(g.bricks, Brick{
			Destructable: false,
			HP:           -1,
			PosX:         brickWidth * 4,
			PosY:         ystart + brickHeight*4,
			Width:        brickWidth,
			Height:       brickHeight,
		})
		g.bricks = append(g.bricks, Brick{
			Destructable: false,
			HP:           -1,
			PosX:         brickWidth * 6,
			PosY:         ystart + brickHeight*4,
			Width:        brickWidth,
			Height:       brickHeight,
		})
		g.bricks = append(g.bricks, Brick{
			Destructable: false,
			HP:           -1,
			PosX:         brickWidth * 8,
			PosY:         ystart + brickHeight*4,
			Width:        brickWidth,
			Height:       brickHeight,
		})
	case 4: // Stage 4
		for l := 0; l < 3; l++ {
			for i := 0; i < 11; i++ {
				g.bricks = append(g.bricks, Brick{
					Destructable: true,
					HP:           2 + 1 - l,
					PosX:         brickWidth * float64(i),
					PosY:         ystart + (brickHeight * float64(l)),
					Width:        brickWidth,
					Height:       brickHeight,
				})

				g.brickCount++
			}
		}
		for l := 0; l < 3; l++ {
			for i := 0; i < 11; i++ {
				g.bricks = append(g.bricks, Brick{
					Destructable: true,
					HP:           1 + l,
					PosX:         brickWidth * float64(i),
					PosY:         ystart + ((brickHeight) * float64(l+4)),
					Width:        brickWidth,
					Height:       brickHeight,
				})

				g.brickCount++
			}
		}
	case 5: // Stage 5
		for l := 0; l < 2; l++ {
			for i := 0; i < 11; i++ {
				g.bricks = append(g.bricks, Brick{
					Destructable: true,
					HP:           2 + 1 - l,
					PosX:         brickWidth * float64(i),
					PosY:         ystart + (brickHeight * float64(l)),
					Width:        brickWidth,
					Height:       brickHeight,
				})

				g.brickCount++
			}
		}
		for l := 0; l < 2; l++ {
			for i := 0; i < 11; i++ {
				g.bricks = append(g.bricks, Brick{
					Destructable: true,
					HP:           2 + 1 - l,
					PosX:         brickWidth * float64(i),
					PosY:         ystart + (brickHeight * float64(l+2)),
					Width:        brickWidth,
					Height:       brickHeight,
				})

				g.brickCount++
			}
		}

		for i := 0; i < 11; i++ {
			if i%2 == 0 {
				g.bricks = append(g.bricks, Brick{
					Destructable: false,
					HP:           -1,
					PosX:         brickWidth * float64(i),
					PosY:         ystart + brickHeight*4,
					Width:        brickWidth,
					Height:       brickHeight,
				})
			}
		}
	}
}
