package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"

	"github.com/mazei513/sugoku/game"
	"github.com/mazei513/sugoku/stringerr"
)

const regularExit stringerr.StringErr = "expected exit"

var gameState *game.Game

func update(screen *ebiten.Image) error {
	if err := gameState.Update(); err != nil {
		return err
	}

	if gameState.ToExit() {
		return regularExit
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	return gameState.Draw(screen)
}

func main() {
	var err error
	gameState, err = game.NewGame()
	if err != nil {
		log.Fatal(err)
	}

	if err := ebiten.Run(update, game.WindowSize, game.WindowSize, 1.0, "Sugoku"); err != nil && err != regularExit {
		log.Fatal(err)
	}
}
