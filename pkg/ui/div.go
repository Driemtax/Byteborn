package ui

import (
	"image/color"

	"github.com/Driemtax/Byteborn/pkg/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Div struct {
	text *Text

	posX   int
	posY   int
	width  int
	height int
}

func NewDiv(width, height int, msg string) *Div {
	return &Div{
		text:   &Text{content: msg, font: util.InitFont(), justificationX: text.AlignCenter, justificationY: text.AlignCenter},
		width:  width,
		height: height,
	}
}

func (d *Div) Update(posX, posY int) {
	// input is the middle of the screen now still need to center the div
	newX := posX - (d.width / 2)
	newY := posY - (d.height / 2)

	d.posX = newX
	d.posY = newY
}

func (d *Div) Draw(screen *ebiten.Image) {
	// Draw a black border for the whole div to make it visible
	vector.FillRect(screen, float32(d.posX), float32(d.posY), float32(d.width), float32(d.height), color.Black, true)

	// Draw text in the middle of the div
	x := (d.width / 2) + d.posX
	y := (d.height / 2) + d.posY

	d.text.Draw(screen, x, y)
}
