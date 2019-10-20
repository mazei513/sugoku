package game

import (
	"github.com/hajimehoshi/ebiten"
)

type board struct {
	squares []*Square

	boardImage *ebiten.Image
	drawOpts   *ebiten.DrawImageOptions
}

func newEmptyBoard() (*board, error) {
	squares := make([]*Square, 0, 9*9)
	for x := uint8(0); x < 9; x++ {
		for y := uint8(0); y < 9; y++ {
			square, err := NewSquare(x%2 == 0, y+1, x, y)
			if err != nil {
				return nil, err
			}
			squares = append(squares, square)
		}
	}

	boardSize := 9*9*squareDrawSize + 2*groupPadSize
	boardImage, err := ebiten.NewImage(boardSize, boardSize, ebiten.FilterNearest)
	if err != nil {
		return nil, err
	}
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(8, 8)

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
