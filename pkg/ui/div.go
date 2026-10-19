package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Div struct {
	components []Component

	posX   int
	posY   int
	width  int
	height int
}

func NewDiv(width, height, posX, posY int) *Div {
	return &Div{
		components: []Component{},
		width:      width,
		height:     height,
		posX:       posX,
		posY:       posY,
	}
}

func (d *Div) AddChild(c Component) {
	d.components = append(d.components, c)
}

func (d *Div) OnKeyPress(key ebiten.Key) {

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
	//vector.FillRect(screen, float32(d.posX), float32(d.posY), float32(d.width), float32(d.height), color.Black, true)
	vector.StrokeRect(screen, float32(d.posX), float32(d.posY), float32(d.width), float32(d.height), 2.0, color.White, false)

	// Draw text in the middle of the div
	//x := (d.width / 2) + d.posX
	//y := (d.height / 2) + d.posY
	x := d.posX
	y := d.posY

	for _, c := range d.components {
		c.Draw(screen, x, y, d.width, d.height)
	}
}
