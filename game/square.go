package game

import (
	"fmt"
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
)

// InvalidValueErr describes an invalid value being set to a square
type InvalidValueErr uint8

func (i InvalidValueErr) Error() string {
	return fmt.Sprintf("invalid value given: %d", i)
}

const squareDrawSize = 64

type square struct {
	isSelected bool
	isEditable bool
	value      uint8
	x, y       uint8
}

func (s square) getImage() (*ebiten.Image, error) {
	outer, err := ebiten.NewImage(squareDrawSize, squareDrawSize, ebiten.FilterNearest)
	if err != nil {
		return nil, err
	}

	outerColor := color.NRGBA{0, 0, 0, 0xFF}
	if s.isSelected {
		outerColor = color.NRGBA{0xCC, 0x22, 0x22, 0xFF}
	}
	outer.Fill(outerColor)

	inner, err := ebiten.NewImage(squareDrawSize-4, squareDrawSize-4, ebiten.FilterNearest)
	inner.Fill(color.White)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(2, 2)

	outer.DrawImage(inner, opts)

	return outer, nil
}

func (s square) xDrawPos(boardOffsetX float64) float64 {
	xOffset := float64(s.x/3)*4 + boardOffsetX
	return float64(s.x)*squareDrawSize + xOffset
}

func (s square) yDrawPos(boardOffsetY float64) float64 {
	yOffset := float64(s.y/3)*4 + boardOffsetY
	return float64(s.y)*squareDrawSize + yOffset
}

func (s square) drawSquare(screen *ebiten.Image, boardOffsetX, boardOffsetY float64) error {
	img, err := s.getImage()
	if err != nil {
		return err
	}
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(s.xDrawPos(boardOffsetX), s.yDrawPos(boardOffsetY))

	screen.DrawImage(img, opts)
	return nil
}

func (s square) xTextPos(boardOffsetX float64) int {
	return int(s.xDrawPos(boardOffsetX)) + 14
}

func (s square) yTextPos(boardOffsetY float64) int {
	return int(s.yDrawPos(boardOffsetY)) + squareDrawSize - 12
}

func (s square) drawValueText(screen *ebiten.Image, boardOffsetX, boardOffsetY float64) {
	textColor := color.NRGBA{0, 0, 0, 0xFF}
	if s.isEditable {
		textColor = color.NRGBA{0x33, 0x33, 0xBB, 0xFF}
	}
	text.Draw(screen, strconv.Itoa(int(s.value)), mplusNormalFont, s.xTextPos(boardOffsetX), s.yTextPos(boardOffsetY), textColor)
}

func (s square) Draw(screen *ebiten.Image, boardOffsetX, boardOffsetY float64) error {
	if err := s.drawSquare(screen, boardOffsetX, boardOffsetY); err != nil {
		return err
	}
	s.drawValueText(screen, boardOffsetX, boardOffsetY)
	return nil
}
