package game

import (
	"log"

	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Game struct {
	windowWidth  int32
	windowHeight int32
	window       *sdl.Window
	renderer     *sdl.Renderer
	font         *ttf.Font
	textures     map[string]*sdl.Texture
	sfx          map[string]*mix.Chunk
	Active       bool
	mouse        Mouse
	lastMouse    Mouse
	clock        Clock
	stage        int
	paddle       Paddle
	points       int
	ball         Ball
	ballCount    int
	bricks       []Brick
	brickCount   int
}

func (g *Game) Init(windowWidth, windowHeight int32) {
	g.windowWidth, g.windowHeight = windowWidth, windowHeight

	/* Initialize SDL and SDL subsystems */
	var err error

	// Init SDL
	g.window, err = sdl.CreateWindow(
		"Breakout",
		sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		windowWidth, windowHeight,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Create Renderer
	g.renderer, err = sdl.CreateRenderer(
		g.window, -1, sdl.RENDERER_ACCELERATED,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize SDL Mixer
	err = mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 1, 2048)
	if err != nil {
		log.Fatal(err)
	}
	mix.Volume(-1, 96)
	g.initSfx()

	// Initialize font renderer
	ttf.Init()
	g.font, _ = ttf.OpenFont("res/font.ttf", 120)
	g.initTextures()

	/* Initialize Game Parameters */
	// Game objects
	g.paddle.Init(float64(windowWidth), float64(windowHeight))
	g.ball.Init(windowHeight, g.paddle.PosX)
	g.stage = 0
	g.initBricks()
	g.ballCount = 3
	g.points = 0

	// Clock
	g.clock.Init()
	g.Active = true
}

func (g *Game) initSfx() {
	g.sfx = make(map[string]*mix.Chunk)

	// Brick break sfx
	g.sfx["break"], _ = mix.LoadWAV("res/break.wav")

	// Wall bounce
	g.sfx["bounce"], _ = mix.LoadWAV("res/bounce.wav")

	// Paddle miss
	g.sfx["miss"], _ = mix.LoadWAV("res/miss.wav")
}

func (g *Game) deleteSfx() {
	for key := range g.sfx {
		g.sfx[key].Free()
	}
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
	g.textures["start"], _ = g.renderer.CreateTextureFromSurface(textSurface)
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
	g.textures["score"], _ = g.renderer.CreateTextureFromSurface(textSurface)
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
	g.textures["balls"], _ = g.renderer.CreateTextureFromSurface(textSurface)
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
	g.textures["game over"], _ = g.renderer.CreateTextureFromSurface(textSurface)
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
	g.textures["restart"], _ = g.renderer.CreateTextureFromSurface(textSurface)
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
	g.textures["breakout"], _ = g.renderer.CreateTextureFromSurface(textSurface)
	textSurface.Free()

	// Congratulations
	textSurface, _ = g.font.RenderUTF8Solid(
		"congratulations",
		sdl.Color{
			R: 255,
			G: 255,
			B: 255,
		},
	)
	g.textures["congratulations"], _ = g.renderer.CreateTextureFromSurface(textSurface)
	textSurface.Free()
}

func (g *Game) deleteTextures() {
	for key := range g.textures {
		g.textures[key].Destroy()
	}
}

func (g *Game) Quit() {
	// SDl Mixer
	g.deleteSfx()
	mix.CloseAudio()

	// Font renderer
	g.font.Close()
	ttf.Quit()

	// Free texture memory
	g.deleteTextures()

	// Window stuff
	g.renderer.Destroy()
	g.window.Destroy()
	sdl.Quit()
}

func (g *Game) setStage(stage int) {
	sdl.Delay(750)
	g.stage = stage
	g.initBricks()

	g.ball.Held = true
	g.ballCount = 3

	if stage == 1 {
		g.points = 0
	}
}

func (g *Game) Update() {
	// Update gamestate
	g.handleInput()

	// Physics and collision checks
	g.updateBall()

	bbc := 0
	for i := range g.bricks {
		g.ball.BrickCollide(&g.bricks[i], &g.points, g.sfx)

		// Count destroyed bricks
		if g.bricks[i].Destructable && g.bricks[i].HP == 0 {
			bbc += 1
		}
	}

	// Draw
	g.draw()

	// Check to see if stage is complete
	// Compare no. of destroyable bricks to
	// current brick count
	/* These are below draw call so you see menu brick destroyed */
	if bbc == g.brickCount {
		switch g.stage {
		case 999:
			g.setStage(1)
		case 6:
			break
		default:
			g.setStage(g.stage + 1)
		}
	}

	if g.stage != 6 && g.ballCount == 0 {
		g.setStage(999)
	}

	// Update delta time
	g.clock.Tick()
}
