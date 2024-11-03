package game

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

type Game struct {
	WindowWidth  int32
	WindowHeight int32
	Window       *sdl.Window
	Renderer     *sdl.Renderer
	Active       bool
	LastMouse    Mouse
	Mouse        Mouse
	Clock        Clock
	Paddle       Paddle
	Ball         Ball
}

func (g *Game) Init(windowWidth, windowHeight int32) {
	g.WindowWidth, g.WindowHeight = windowWidth, windowHeight

	/* Initialize SDL and SDL subsystems */
	var err error

	// Init SDL
	g.Window, err = sdl.CreateWindow(
		"Breakout",
		sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		windowWidth, windowHeight,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Create Renderer
	g.Renderer, err = sdl.CreateRenderer(
		g.Window, -1, sdl.RENDERER_ACCELERATED,
	)
	if err != nil {
		log.Fatal(err)
	}

	/* Initialize Game Parameters */
	// Game objects
	g.Paddle.Init(windowWidth, windowHeight)
	g.Ball.Init(windowHeight, g.Paddle.PosX)

	// Clock
	g.Clock.Init()
	g.Active = true
}

func (g *Game) SetWindowSize(windowWidth, windowHeight int32) {
	g.Window.SetSize(
		windowWidth, windowHeight,
	)

	g.Paddle.Init(windowWidth, windowHeight)
}

func (g *Game) Quit() {
	g.Renderer.Destroy()
	g.Window.Destroy()
	sdl.Quit()
}

func (g *Game) Update() {
	// Update gamestate
	g.HandleInput()
	//g.Ball.Update(g.WindowWidth, g.WindowHeight, g.Paddle, g.Clock.DeltaTime)
	g.updateBall()

	// Draw
	g.Draw()

	// Update delta time
	g.Clock.Tick()
}
