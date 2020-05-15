package game

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
)

const (
	boardSize = 9*squareDrawSize + 2*groupPadSize
	boardPad  = 8
)

var selImg *ebiten.Image

func init() {
	const innerPad = 4
	const innerSize = squareDrawSize - innerPad*2
	inner, _ := ebiten.NewImage(innerSize, innerSize, ebiten.FilterDefault)
	selImg, _ = ebiten.NewImage(squareDrawSize, squareDrawSize, ebiten.FilterDefault)

	inner.Fill(color.White)
	selImg.Fill(color.RGBA{0xff, 0, 0, 0xff})

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(innerPad, innerPad)
	selImg.DrawImage(inner, opts)
}

type board struct {
	squares []*Square

	boardImage *ebiten.Image
	drawOpts   *ebiten.DrawImageOptions
}

func newBoardFromString(boardString string) (*board, error) {
	boardImage, _ := ebiten.NewImage(boardSize, boardSize, ebiten.FilterNearest)
	boardImage.Fill(color.Black)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(boardPad, boardPad)

	squares := make([]*Square, 0, 9*9)
	for y := uint8(0); y < 9; y++ {
		for x := uint8(0); x < 9; x++ {
			value := boardString[y*9+x] - '0'
			square := NewSquare(value == 0, value, x, y)
			boardImage.DrawImage(square.image, square.drawOpts)
			if !square.isEditable {
				text.Draw(boardImage, strconv.Itoa(int(value)), squareNumberFont, square.textPosX, square.textPosY, square.textColor())
			}
			squares = append(squares, square)
		}
	}

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
		screen.DrawImage(b.boardImage, b.drawOpts)
		for _, s := range b.squares {
			if s.isSelected {
				opts := &ebiten.DrawImageOptions{}
				opts.GeoM.Translate(float64(s.rect.Min.X)+boardPad, float64(s.rect.Min.Y)+boardPad)
				screen.DrawImage(selImg, opts)
			}
			if s.isEditable && s.value != 0 || !s.isEditable && s.isDuplicate {
				text.Draw(screen, strconv.Itoa(int(s.value)), squareNumberFont, s.textPosX+boardPad, s.textPosY+boardPad, s.textColor())
			}
		}
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
