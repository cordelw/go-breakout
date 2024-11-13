package main

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/cordelw/go-breakout/game"
)

func loadCFG() (int32, int32) {
	cfgFile, _ := os.ReadFile("res/cfg.json")
	jsonData := make(map[string]string)
	json.Unmarshal(cfgFile, &jsonData)

	w, _ := strconv.Atoi(jsonData["window_width"])
	h, _ := strconv.Atoi(jsonData["window_height"])

	return int32(w), int32(h)
}

func main() {
	ww, wh := loadCFG()
	game := new(game.Game)

	if ww != 0 && wh != 0 {
		game.Init(ww, wh)
	} else {
		game.Init(320, 240)
	}
	defer game.Quit()

	for game.Active {
		game.Update()
	}
}
