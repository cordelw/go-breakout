package game

import (
	"math"
	"math/rand"

	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

type Ball struct {
	PosX, PosY float64
	VelX, VelY float64
	Speed      float64
	Radius     int
	Held       bool
}

func (b *Ball) Init(windowHeight int32, posX float64) {
	b.Radius = int(windowHeight) / 48
	b.PosY = float64(windowHeight-(windowHeight/8)) - float64(b.Radius+b.Radius/2)
	b.PosX = posX
	b.Speed = 0.3

	b.Held = true
}

func (b Ball) Draw(renderer *sdl.Renderer) {
	renderer.SetDrawColor(255, 255, 255, 255)
	DrawCircle(renderer, int(b.PosX), int(b.PosY), b.Radius)
}

func (g *Game) updateBall() {
	b := &g.ball

	// Held by player before being released
	// Follow player paddle
	if b.Held {
		b.PosY = float64(g.windowHeight-(g.windowHeight/8)) - float64(b.Radius+b.Radius/2)
		b.PosX = g.paddle.PosX + float64(g.paddle.Width/2)
		//b.PosX = float64(g.Mouse.PosX)
		//b.PosY = float64(g.Mouse.PosY)
		return
	}

	// Fix Horizontal Velocity
	if b.VelX > b.Speed {
		b.VelX = b.Speed
	} else if b.VelX < -b.Speed {
		b.VelX = -b.Speed
	}

	// Fix Vertical Velocity
	if b.VelY > b.Speed/2 {
		b.VelY = b.Speed / 2
	} else if b.VelY < -b.Speed/2 {
		b.VelY = -b.Speed / 2
	}

	// Adjust velocity for different
	// screen sizes
	h := float64(g.windowWidth) / 440
	v := float64(g.windowHeight) / 330

	// Update position
	b.PosX += b.VelX * h * g.clock.DeltaTime
	b.PosY += b.VelY * v * g.clock.DeltaTime

	// Bounds collisions
	// Left
	if b.PosX-float64(b.Radius) < 0 {
		b.PosX = 0 + float64(b.Radius)
		b.VelX = -b.VelX
		g.sfx["bounce"].Play(-1, 0)

		// Right
	} else if b.PosX+float64(b.Radius) > float64(g.windowWidth) {
		b.PosX = float64(g.windowWidth - int32(b.Radius))
		b.VelX = -b.VelX
		g.sfx["bounce"].Play(-1, 0)
	}

	// Top
	if b.PosY-float64(b.Radius) < 0 {
		b.PosY = 0 + float64(b.Radius)
		b.VelY = -b.VelY
		g.sfx["bounce"].Play(-1, 0)

		// Bottom
	} else if b.PosY > float64(g.windowHeight) {
		g.ballCount -= 1
		b.Held = true
		g.sfx["miss"].Play(-1, 0)
	}

	// Player paddle collisions
	// Check horizontal

	if b.PosX >= g.paddle.PosX && b.PosX <= g.paddle.PosX+float64(g.paddle.Width) {
		// Check vertical
		if b.PosY+float64(b.Radius) > g.paddle.PosY && b.PosY < g.paddle.PosY {
			// Reposition Ball and reverse Y direction
			b.PosY = g.paddle.PosY - float64(b.Radius)
			b.VelY = -b.VelY

			// Determine which direction the paddle is
			// moving defined as -1, 0, or 1
			var dirNorm float64
			mDiff := g.mouse.PosX - g.lastMouse.PosX
			if mDiff > 0 {
				dirNorm = 1
			} else if mDiff < 0 {
				dirNorm = -1
			} else {
				dirNorm = 0
			}

			// Set velocity to random variance between
			// 0.1 and b.Speed
			var newVelX float64
			newVelX = randFloatN(10, int(b.Speed*100))

			// Multiplied by the direction if not 0
			if dirNorm != 0 {
				newVelX *= dirNorm
			}

			b.VelX = newVelX

			g.sfx["bounce"].Play(-1, 0)
		}
	}
}

func randFloatN(min, max int) float64 {
	return float64(rand.Intn(max-min)+min) / 100
}

func (b *Ball) BrickCollide(brick *Brick, score *int, sfx map[string]*mix.Chunk) {
	/*
		Clamp function and sh:
		https://www.youtube.com/watch?v=_xj8FyG-aac

		Collision Resolution:
		https://www.youtube.com/watch?v=be0WANYMH_k
	*/

	// Do not check dead bricks
	if brick.HP == 0 {
		return
	}

	// Only run precise collision checks if ball is inside the
	// horizontal bounds of the brick
	if b.PosX-float64(b.Radius) > brick.PosX+brick.Width || b.PosX+float64(b.Radius) < brick.PosX {
		return
	}

	// Proper collision check
	nearestX := clamp(brick.PosX, brick.PosX+brick.Width, b.PosX)
	nearestY := clamp(brick.PosY, brick.PosY+brick.Height, b.PosY)

	distX := b.PosX - nearestX
	distY := b.PosY - nearestY
	dist := (distX * distX) + (distY * distY)

	// Do nothing if distance is greater than
	// circle radius
	if dist > float64(b.Radius*b.Radius) {
		return
	}

	// Collision response
	distSqrt := math.Sqrt(dist)
	normX := distX / distSqrt
	normY := distY / distSqrt

	// Reverse velocities accordingly
	// Prevent buggy reversion by checking to see if
	// the normals have the same sign as the velocity
	// before reversing
	if normX != 0 && (b.VelX < 0) != (normX < 0) {
		b.VelX = -b.VelX
	}
	if normY != 0 && (b.VelY < 0) != (normY < 0) {
		b.VelY = -b.VelY
	}

	// Reposition Ball to where it hits
	// the brick
	if distX != 0 {
		b.PosX = nearestX + (float64(b.Radius) * normX)
	}
	if distY != 0 {
		b.PosY = nearestY + (float64(b.Radius) * normY)
	}

	// Damage brick if destructable
	if brick.Destructable {
		brick.HP -= 1
		*score += 100

		if brick.HP == 0 {
			sfx["break"].Play(-1, 0)
		} else {
			sfx["bounce"].Play(-1, 0)
		}
	} else {
		sfx["bounce"].Play(-1, 0)
	}
}

func clamp(min, max, value float64) float64 {
	return math.Max(min, math.Min(max, value))
}
