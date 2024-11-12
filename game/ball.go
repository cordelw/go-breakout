package game

import (
	"math"

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

func (b *Ball) BrickCollide(dt float64, brick *Brick) {
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
	radSqrt := math.Sqrt(dist)
	normX := distX / radSqrt
	normY := distY / radSqrt

	// Reverse velocities accordingly
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
	}
}

func clamp(min, max, value float64) float64 {
	return math.Max(min, math.Min(max, value))
}

func (g *Game) updateBall() {
	b := &g.Ball

	// Held by player before being released
	// Follow player paddle
	if b.Held {
		b.PosY = float64(g.WindowHeight-(g.WindowHeight/8)) - float64(b.Radius+b.Radius/2)
		b.PosX = g.Paddle.PosX + float64(g.Paddle.Width/2)
		//b.PosX = float64(g.Mouse.PosX)
		//b.PosY = float64(g.Mouse.PosY)
		return
	}

	// Fix Horizontal Velocity
	if b.VelX > b.Speed {
		b.VelX = b.Speed
	}

	// Update position
	b.PosX += b.VelX * g.Clock.DeltaTime
	b.PosY += b.VelY * g.Clock.DeltaTime

	// Bounds collisions
	// Left
	if b.PosX-float64(b.Radius) < 0 {
		b.PosX = 0 + float64(b.Radius)
		b.VelX = -b.VelX
	}
	// Right
	if b.PosX+float64(b.Radius) > float64(g.WindowWidth) {
		b.PosX = float64(g.WindowWidth - int32(b.Radius))
		b.VelX = -b.VelX
	}
	// Top
	if b.PosY-float64(b.Radius) < 0 {
		b.PosY = 0 + float64(b.Radius)
		b.VelY = -b.VelY
	}
	// Bottom
	if b.PosY > float64(g.WindowHeight) {
		b.Held = true
	}

	// Player paddle collisions
	// Check horizontal

	if b.PosX >= g.Paddle.PosX && b.PosX <= g.Paddle.PosX+float64(g.Paddle.Width) {
		// Check vertical
		if b.PosY+float64(b.Radius) > g.Paddle.PosY && b.PosY < g.Paddle.PosY {
			b.PosY = g.Paddle.PosY - float64(b.Radius)
			b.VelY = -b.VelY

			dir := float64(g.LastMouse.PosX - g.Mouse.PosX)
			if dir != 0 {
				b.VelX = -(math.Min(dir*10, 0.3))
			}
		}
	}
}
