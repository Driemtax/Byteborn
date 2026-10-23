package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Triangle struct {
	img *ebiten.Image
}

func NewTriangle(width, height int) *Triangle {
	new_img := ebiten.NewImage(width, height)
	actualWidth := width

	// width need to be an odd number to draw a rectangle otherwise the tip is missing
	if width%2 == 0 {
		actualWidth += 1
	}

	p := &vector.Path{}
	p.MoveTo(0, 0)
	p.LineTo(float32(actualWidth), float32(height/2))
	p.LineTo(0, float32(height))
	p.Close()

	vector.FillPath(new_img, p, &vector.FillOptions{}, &vector.DrawPathOptions{
		AntiAlias:  true,
		ColorScale: ebiten.ColorScale{},
	})
	return &Triangle{
		img: new_img,
	}
}

func (t *Triangle) Draw(screen *ebiten.Image, posX int, posY int, width int, height int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(posX), float64(posY))
	screen.DrawImage(t.img, op)
}

// Will not be called
func (t *Triangle) OnKeyPress(ebiten.Key) {

}

// Will not be called
func (t *Triangle) GetSelectedButtonLabel() string {
	return ""
}
