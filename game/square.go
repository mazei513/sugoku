package game

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

// InvalidValueErr describes an invalid value being set to a square
type InvalidValueErr uint8

func (i InvalidValueErr) Error() string {
	return fmt.Sprintf("invalid value given: %d", i)
}

const (
	squareDrawSize = 64
	groupPadSize   = 12
)

// Square handles a single value of a sudoku square
type Square struct {
	// mutable state
	isSelected  bool
	value       uint8
	isDuplicate bool

	// constant values
	isEditable         bool
	x, y               uint8
	rect               image.Rectangle
	textPosX, textPosY int
	image              *ebiten.Image
	drawOpts           *ebiten.DrawImageOptions
}

func squarePos(v uint8) int {
	return int(v)*squareDrawSize + int(v)/3*groupPadSize
}

// NewSquare creates a new square with the given params
func NewSquare(isEditable bool, value, x, y uint8) *Square {
	posX := squarePos(x)
	posY := squarePos(y)

	const innerOffset = 4
	const innerSize = squareDrawSize - innerOffset*2
	img, _ := ebiten.NewImage(innerSize, innerSize, ebiten.FilterDefault)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(posX)+innerOffset, float64(posY)+innerOffset)

	if isEditable {
		img.Fill(color.White)
	} else {
		img.Fill(color.Gray{0xa0})
	}

	return &Square{
		isSelected:  false,
		isDuplicate: false,
		isEditable:  isEditable,
		value:       value,
		x:           x,
		y:           y,
		rect:        image.Rect(posX, posY, posX+squareDrawSize, posY+squareDrawSize),
		textPosX:    posX + 14,
		textPosY:    posY + squareDrawSize - 12,
		image:       img,
		drawOpts:    opts,
	}
}

// Update handles the square's state
func (s *Square) Update() {
	if s.isEditable {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			s.isSelected = false
			x, y := ebiten.CursorPosition()
			if s.rect.Min.X <= x && x < s.rect.Max.X && s.rect.Min.Y <= y && y < s.rect.Max.Y {
				s.isSelected = true
			}
		}
		if s.isSelected {
			if ebiten.IsKeyPressed(ebiten.Key0) || ebiten.IsKeyPressed(ebiten.KeyBackspace) {
				s.value = 0
			}

			keys := []ebiten.Key{1, 2, 3, 4, 5, 6, 7, 8, 9}
			for _, k := range keys {
				if ebiten.IsKeyPressed(k) {
					s.value = uint8(k)
				}
			}
		}
	}
}

func (s Square) outerColor() color.Color {
	if s.isSelected {
		return color.NRGBA{0xCC, 0x22, 0x22, 0xFF}
	}
	return color.NRGBA{0, 0, 0, 0xFF}
}

func (s Square) innerColor() color.Color {
	if s.isEditable {
		return color.White
	}
	return color.RGBA{0xa0, 0xa0, 0xa0, 0xff}
}

func (s Square) textColor() color.Color {
	if s.isEditable {
		return color.NRGBA{0x33, 0x33, 0xBB, 0xFF}
	} else if s.isDuplicate {
		return color.RGBA{255, 0, 0, 255}
	}
	return color.NRGBA{0, 0, 0, 0xFF}
}
