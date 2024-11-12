package main

import (
	"github.com/cordelw/go-breakout/game"
)

func main() {
	game := new(game.Game)
	game.Init(400*2, 300*2)
	defer game.Quit()

	for game.Active {
		game.Update()
	}
}
