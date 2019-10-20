package game

import (
	"github.com/hajimehoshi/ebiten"
)

type board struct {
	squares []*square
}

func newEmptyBoard() *board {
	squares := make([]*square, 0, 9*9)
	for x := uint8(0); x < 9; x++ {
		for y := uint8(0); y < 9; y++ {
			squares = append(squares, &square{value: y + 1, isEditable: x%2 == 0, x: x, y: y})
		}
	}
	squares[24].isSelected = true
	return &board{
		squares: squares,
	}
}

func (b board) Draw(screen *ebiten.Image) error {
	for _, s := range b.squares {
		s.Draw(screen, 4, 4)
	}
	return nil
}
