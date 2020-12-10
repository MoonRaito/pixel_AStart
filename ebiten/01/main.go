package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
	"log"
	"pixel_AStart/ebiten/common"
	"pixel_AStart/ebiten/start"
)

func main() {
	game, err := start.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	// Sepcify the window size as you like. Here, a doulbed size is specified.
	ebiten.SetWindowSize(common.ScreenWidth, common.ScreenHeight)
	ebiten.SetWindowTitle("Your game's title")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
