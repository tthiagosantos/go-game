package main

import (
	"game/game"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {

	g := game.NewGame()

	err := ebiten.RunGame(g)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

}
