package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

const (
	// WindowSize is the minimum size the window must be to fit the board
	WindowSize = boardSize + 2*boardPad
)

// Game controls the game state
type Game struct {
	toExit bool
	board  *board
	debug  bool
}

// NewGame creates a new instance of Game
func NewGame() (*Game, error) {
	// board, err := newEmptyBoard()
	board, err := newBoardFromString(`600008517012000000007601824500920603260003040003105200100480762906050300020010050`)
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
	} else if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		g.debug = !g.debug
	}

	g.board.Update()

	return nil
}

// Draw draws the current game state
func (g Game) Draw(screen *ebiten.Image) error {
	g.board.Draw(screen)
	if g.debug {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("%f %f", ebiten.CurrentFPS(), ebiten.CurrentTPS()))
	}
	return nil
}
