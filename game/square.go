package game

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
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

var (
	innerSquareImage       *ebiten.Image
	innerSquareDrawOptions = &ebiten.DrawImageOptions{}
	outerColorUnselected   = color.NRGBA{0, 0, 0, 0xFF}
	outerColorSelected     = color.NRGBA{0xCC, 0x22, 0x22, 0xFF}
	textColorUneditable    = color.NRGBA{0, 0, 0, 0xFF}
	textColorEditable      = color.NRGBA{0x33, 0x33, 0xBB, 0xFF}
)

func init() {
	var err error
	innerSquareImage, err = ebiten.NewImage(squareDrawSize-8, squareDrawSize-8, ebiten.FilterNearest)
	if err != nil {
		log.Fatal(err)
	}
	innerSquareImage.Fill(color.White)
	innerSquareDrawOptions.GeoM.Translate(4, 4)
}

// Square handles a single value of a sudoku square
type Square struct {
	// mutable state
	isSelected bool
	value      uint8

	// constant values
	isEditable         bool
	x, y               uint8
	rect               image.Rectangle
	textPosX, textPosY int
	outerImage         *ebiten.Image
	drawOpts           *ebiten.DrawImageOptions
	textColor          color.Color
}

func minPos(v int) int {
	return v*squareDrawSize + v/3*groupPadSize
}

// NewSquare creates a new square with the given params
func NewSquare(isEditable bool, value, x, y uint8) (*Square, error) {
	minX := minPos(int(x))
	minY := minPos(int(y))

	outer, err := ebiten.NewImage(squareDrawSize, squareDrawSize, ebiten.FilterNearest)
	if err != nil {
		return nil, err
	}
	drawOpts := &ebiten.DrawImageOptions{}
	drawOpts.GeoM.Translate(float64(minX), float64(minY))

	textColor := textColorUneditable
	if isEditable {
		textColor = textColorEditable
	}

	return &Square{
		isSelected: false,
		isEditable: isEditable,
		value:      value,
		x:          x,
		y:          y,
		rect:       image.Rect(minX, minY, minX+squareDrawSize, minY+squareDrawSize),
		textPosX:   minX + 14,
		textPosY:   minY + squareDrawSize - 12,
		outerImage: outer,
		drawOpts:   drawOpts,
		textColor:  textColor,
	}, nil
}

// Update handles the square's state
func (s *Square) Update() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		s.isSelected = false
		x, y := ebiten.CursorPosition()
		if s.rect.Min.X <= x && x < s.rect.Max.X && s.rect.Min.Y <= y && y < s.rect.Max.Y {
			s.isSelected = true
		}
	}
}

func (s Square) drawSquare(screen *ebiten.Image) {
	outerColor := outerColorUnselected
	if s.isSelected {
		outerColor = outerColorSelected
	}
	s.outerImage.Fill(outerColor)
	s.outerImage.DrawImage(innerSquareImage, innerSquareDrawOptions)

	screen.DrawImage(s.outerImage, s.drawOpts)
}

func (s Square) drawValueText(screen *ebiten.Image) {
	text.Draw(screen, strconv.Itoa(int(s.value)), squareNumberFont, s.textPosX, s.textPosY, s.textColor)
}

// Draw draws the square
func (s Square) Draw(screen *ebiten.Image) error {
	s.drawSquare(screen)
	s.drawValueText(screen)
	return nil
}
