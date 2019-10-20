package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// Game controls the game state
type Game struct {
	toExit bool
}

// NewGame creates a new instance of Game
func NewGame() *Game {
	return &Game{}
}

// ToExit signals the program to exit
func (g Game) ToExit() bool {
	return g.toExit
}

// HandleInput updates the game state based on the current input
func (g *Game) HandleInput() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		fmt.Println("here")
		g.toExit = true
	}

	return nil
}
