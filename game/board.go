package game

import (
	"image/color"
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
		b.markDuplicates(s)
	}
}

func (b board) Draw(screen *ebiten.Image) error {
	if b.isComplete() {
		screen.Fill(color.RGBA{0, 255, 0, 255})
	} else {
		for _, s := range b.squares {
			if err := s.Draw(b.boardImage); err != nil {
				return err
			}
		}
		screen.DrawImage(b.boardImage, b.drawOpts)
	}
	return nil
}

func (b board) squareAt(x, y uint8) *Square {
	return b.squares[y*9+x]
}

func (b board) markDuplicates(s *Square) {
	s.isDuplicate = false

	// check in row
	for x := uint8(0); x < 9; x++ {
		if x == s.x {
			continue
		}
		if b.squareAt(x, s.y).value == s.value {
			b.squareAt(x, s.y).isDuplicate = true
			s.isDuplicate = true
		}
	}

	// check in column
	for y := uint8(0); y < 9; y++ {
		if y == s.y {
			continue
		}
		if b.squareAt(s.x, y).value == s.value {
			b.squareAt(s.x, y).isDuplicate = true
			s.isDuplicate = true
		}
	}

	// check in box
	yBox := s.y / 3 * 3
	xBox := s.x / 3 * 3
	for y := yBox; y < yBox+3; y++ {
		for x := xBox; x < xBox+3; x++ {
			if x == s.x && y == s.y {
				continue
			}
			if b.squareAt(x, y).value == s.value {
				b.squareAt(x, y).isDuplicate = true
				s.isDuplicate = true
			}
		}
	}
}

func (b board) isComplete() bool {
	for _, s := range b.squares {
		if s.value == 0 || s.isDuplicate {
			return false
		}
	}
	return true
}
