package ui

import (
	"github.com/Driemtax/Byteborn/pkg/util"
	"github.com/hajimehoshi/ebiten/v2"
)

type Picture struct {
	file   *ebiten.Image
	width  int
	height int
}

func NewPicture(path string) *Picture {
	img := util.LoadAsset(path)
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	return &Picture{
		file:   img,
		width:  width,
		height: height,
	}
}

func (img *Picture) Draw(screen *ebiten.Image, posX int, posY, width int, height int) {
	op := &ebiten.DrawImageOptions{}
	midX := posX + (width / 2) - (img.width / 2)
	midY := posY + (height / 2) - (img.height / 2)
	op.GeoM.Translate(float64(midX), float64(midY))

	screen.DrawImage(img.file, op)
}

// Will not be called
func (img *Picture) OnKeyPress(ebiten.Key) {

}

// Will not be called
func (img *Picture) GetSelectedButtonLabel() string {
	return ""
}
