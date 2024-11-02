package game

import "github.com/veandco/go-sdl2/sdl"

type Clock struct {
	TickTime     uint64
	LastTickTime uint64
	DeltaTime    float64
}

func (c *Clock) Init() {
	c.TickTime = sdl.GetPerformanceCounter()
	c.DeltaTime = 0.000001
}

func (c *Clock) Tick() {
	c.LastTickTime = c.TickTime
	c.TickTime = sdl.GetPerformanceCounter()

	c.DeltaTime = (float64(c.TickTime-c.LastTickTime) * 1000) / float64(sdl.GetPerformanceFrequency())
}
