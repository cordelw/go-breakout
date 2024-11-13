package game

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Game struct {
	WindowWidth  int32
	WindowHeight int32
	Window       *sdl.Window
	Renderer     *sdl.Renderer
	font         *ttf.Font
	textures     map[string]*sdl.Texture
	Active       bool
	Mouse        Mouse
	LastMouse    Mouse
	Clock        Clock
	Stage        int
	Paddle       Paddle
	points       int
	Ball         Ball
	ballCount    int
	Bricks       []Brick
	brickCount   int
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

	// Initialize font renderer
	ttf.Init()
	g.font, _ = ttf.OpenFont("font.ttf", 120)
	g.initTextures()

	/* Initialize Game Parameters */
	// Game objects
	g.Paddle.Init(float64(windowWidth), float64(windowHeight))
	g.Ball.Init(windowHeight, g.Paddle.PosX)
	g.Stage = 0
	g.InitBricks()
	g.ballCount = 3
	g.points = 0

	// Clock
	g.Clock.Init()
	g.Active = true
}

func (g *Game) initTextures() {
	g.textures = make(map[string]*sdl.Texture)

	// Start text
	textSurface, _ := g.font.RenderUTF8Solid(
		"Start",
		sdl.Color{
			R: 0,
			G: 0,
			B: 0,
		},
	)
	g.textures["start"], _ = g.Renderer.CreateTextureFromSurface(textSurface)
	textSurface.Free()

	// Score
	textSurface, _ = g.font.RenderUTF8Solid(
		"Score:",
		sdl.Color{
			R: 255,
			G: 255,
			B: 255,
		},
	)
	g.textures["score"], _ = g.Renderer.CreateTextureFromSurface(textSurface)
	textSurface.Free()

	// Balls
	textSurface, _ = g.font.RenderUTF8Solid(
		"Balls:",
		sdl.Color{
			R: 255,
			G: 255,
			B: 255,
		},
	)
	g.textures["balls"], _ = g.Renderer.CreateTextureFromSurface(textSurface)
	textSurface.Free()

	// Game over
	textSurface, _ = g.font.RenderUTF8Solid(
		"Game Over.",
		sdl.Color{
			R: 255,
			G: 255,
			B: 255,
		},
	)
	g.textures["game over"], _ = g.Renderer.CreateTextureFromSurface(textSurface)
	textSurface.Free()

	// Restart
	textSurface, _ = g.font.RenderUTF8Solid(
		"Restart",
		sdl.Color{
			R: 0,
			G: 0,
			B: 0,
		},
	)
	g.textures["restart"], _ = g.Renderer.CreateTextureFromSurface(textSurface)
	textSurface.Free()

	// BREAKOUT
	textSurface, _ = g.font.RenderUTF8Solid(
		"breakout",
		sdl.Color{
			R: 255,
			G: 255,
			B: 255,
		},
	)
	g.textures["breakout"], _ = g.Renderer.CreateTextureFromSurface(textSurface)
	textSurface.Free()
}

func (g *Game) deleteTextures() {
	for key := range g.textures {
		g.textures[key].Destroy()
	}
}

func (g *Game) SetWindowSize(windowWidth, windowHeight int32) {
	g.Window.SetSize(
		windowWidth, windowHeight,
	)

	g.Paddle.Init(float64(windowWidth), float64(windowHeight))
}

func (g *Game) Quit() {
	// Font renderer
	g.font.Close()
	ttf.Quit()

	// Free texture memory
	g.deleteTextures()

	// Window stuff
	g.Renderer.Destroy()
	g.Window.Destroy()
	sdl.Quit()
}

func (g *Game) setStage(stage int) {
	sdl.Delay(750)
	g.Stage = stage
	g.InitBricks()

	g.Ball.Held = true
	g.ballCount = 3

	if stage == 1 {
		g.points = 0
	}
}

func (g *Game) Update() {
	// Update gamestate
	g.HandleInput()

	// Physics and collision checks
	g.updateBall()

	bbc := 0
	for i := range g.Bricks {
		g.Ball.BrickCollide(&g.Bricks[i], &g.points)

		// Count destroyed bricks
		if g.Bricks[i].Destructable && g.Bricks[i].HP == 0 {
			bbc += 1
		}
	}

	// Draw
	g.Draw()

	// Check to see if stage is complete
	// Compare no of destroyable bricks to
	// current brick count
	/* These are below draw call so you see menu brick destroyed */
	if bbc == g.brickCount {
		if g.Stage == 999 {
			g.setStage(1)
		} else {
			g.setStage(g.Stage + 1)
		}
	}

	if g.ballCount == 0 {
		g.setStage(999)
	}

	// Update delta time
	g.Clock.Tick()
}
