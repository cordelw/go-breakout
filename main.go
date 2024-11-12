package main

import (
	"github.com/cordelw/go-breakout/game"
)

func main() {
	game := new(game.Game)
	game.Init(320, 240)
	defer game.Quit()

	for game.Active {
		game.Update()
	}
}
