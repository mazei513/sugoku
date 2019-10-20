package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// Game controls the game state
type Game struct {
	toExit bool
	board  *board
}

// NewGame creates a new instance of Game
func NewGame() (*Game, error) {
	board, err := newEmptyBoard()
	if err != nil {
		return nil, err
	}
	return &Game{board: board}, nil
}

// ToExit signals the program to exit
func (g Game) ToExit() bool {
	return g.toExit
}

// Update updates the game state based on the current input
func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		g.toExit = true
	}

	g.board.Update()

	return nil
}

// Draw draws the current game state
func (g Game) Draw(screen *ebiten.Image) error {
	screen.Fill(color.NRGBA{0, 0, 0, 0xff})

	g.board.Draw(screen)

	return nil
}
