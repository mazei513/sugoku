package game

import (
	"strings"

	"github.com/hajimehoshi/ebiten"
)

const (
	boardSize = 9*squareDrawSize + 2*groupPadSize
	boardPad  = 8
)

type board struct {
	squares []*Square

	boardImage *ebiten.Image
	drawOpts   *ebiten.DrawImageOptions
}

func newEmptyBoard() (*board, error) {
	return newBoardFromString(strings.ReplaceAll(`
000000000
000000000
000000000
000000000
000000000
000000000
000000000
000000000
000000000
`, "\n", ""))
}

func newBoardFromString(boardString string) (*board, error) {
	squares := make([]*Square, 0, 9*9)
	for y := uint8(0); y < 9; y++ {
		for x := uint8(0); x < 9; x++ {
			value := boardString[y*9+x] - '0'
			square, err := NewSquare(value == 0, value, x, y)
			if err != nil {
				return nil, err
			}
			squares = append(squares, square)
		}
	}

	boardImage, err := ebiten.NewImage(boardSize, boardSize, ebiten.FilterNearest)
	if err != nil {
		return nil, err
	}
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(boardPad, boardPad)

	return &board{
		squares:    squares,
		boardImage: boardImage,
		drawOpts:   opts,
	}, nil
}

func (b board) Update() {
	for _, s := range b.squares {
		s.Update()
	}
}

func (b board) Draw(screen *ebiten.Image) error {
	for _, s := range b.squares {
		if err := s.Draw(b.boardImage); err != nil {
			return err
		}
	}
	screen.DrawImage(b.boardImage, b.drawOpts)
	return nil
}
