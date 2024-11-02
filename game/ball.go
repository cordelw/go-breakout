package game

import "github.com/veandco/go-sdl2/sdl"

type Ball struct {
	PosX, PosY float64
	VelX, VelY float64
	Radius     int
	Held       bool
}

func (b *Ball) Init(windowHeight int32, posX float64) {
	b.Radius = int(windowHeight) / 48
	b.PosY = float64(windowHeight-(windowHeight/8)) - float64(b.Radius+b.Radius/2)
	b.PosX = posX

	b.Held = true
}

func (b *Ball) Update(windowWidth, windowHeight int32, paddle Paddle, dt float64) {
	// Held by player before being released
	// Follow player paddle
	if b.Held {
		b.PosX = paddle.PosX + float64(paddle.Width/2)
		return
	}

	// Physics update
	b.PosX += float64(b.VelX) * dt
	b.PosY += float64(b.VelY) * dt

	// Bounds collision detection
	// Left Wall
	if b.PosX-float64(b.Radius) < 0 {
		b.VelX = -b.VelX
		b.PosX = 0 + float64(b.Radius)
	} else if b.PosX+float64(b.Radius) > float64(windowWidth) {
		b.VelX = -b.VelX
		b.PosX = float64(windowWidth - int32(b.Radius))
	}

}

func (b Ball) Draw(renderer *sdl.Renderer) {
	renderer.SetDrawColor(255, 255, 255, 255)
	DrawCircle(renderer, int(b.PosX), int(b.PosY), b.Radius)
}
