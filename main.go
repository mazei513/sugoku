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
	if err := gameState.HandleInput(); err != nil {
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
	gameState = game.NewGame()

	if err := ebiten.Run(update, 640, 640, 1.0, "Sugoku"); err != nil && err != regularExit {
		log.Fatal(err)
	}
}
