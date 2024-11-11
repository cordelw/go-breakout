package main

import (
	"github.com/cordelw/go-breakout/game"
)

func main() {
	game := new(game.Game)
	game.Init(320*2, 240*2)
	defer game.Quit()

	for game.Active {
		game.Update()
	}
}
