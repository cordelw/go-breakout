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

func (b *Ball) RectCollide(dt, rPosX, rPosY, rWidth, rHeight float64) {
	// Get next ball position, then check for collisions based on this,
	// Clamp positions, and then reverse

	/*
		var testPosX, testPosY float64
		nextPosX := b.PosX + (b.VelX * dt)
		nextPosY := b.PosY + (b.VelY * dt)
	*/
}

func (g *Game) updateBall() {
	b := &g.Ball

	// Held by player before being released
	// Follow player paddle
	if b.Held {
		b.PosY = float64(g.WindowHeight-(g.WindowHeight/8)) - float64(b.Radius+b.Radius/2)
		b.PosX = g.Paddle.PosX + float64(g.Paddle.Width/2)
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
				b.VelX = -math.Min(dir*10, 0.3)
			}
		}
	}

	// Updated paddle collide
	//p := g.Paddle
	//b.CheckCollision(p.PosX, p.PosY, p.Width, p.Height)
}
