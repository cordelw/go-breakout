package game

import "github.com/veandco/go-sdl2/sdl"

type Paddle struct {
	PosX, PosY    float64
	Width, Height float64
}

func (p *Paddle) Init(windowWidth, windowHeight float64) {
	p.Width = windowWidth / 11
	p.Height = windowHeight / 24
	p.PosY = float64(windowHeight - (windowHeight / 8))
}

func (p Paddle) Draw(renderer *sdl.Renderer) {
	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.DrawRect(&sdl.Rect{
		X: int32(p.PosX),
		Y: int32(p.PosY),
		W: int32(p.Width),
		H: int32(p.Height),
	})
}
