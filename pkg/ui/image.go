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
	imgWidth := img.file.Bounds().Dx()
	imgHeight := img.file.Bounds().Dy()
	midX := posX + (width / 2) - (imgWidth + 15)
	midY := posY + (height / 2) - (imgHeight)
	op.GeoM.Translate(float64(midX), float64(midY))
	op.GeoM.Scale(2.0, 2.0)

	screen.DrawImage(img.file, op)
	// Circle for debugging the centering
	//vector.FillCircle(screen, float32(midX), float32(midY), 10.0, color.White, false)
}

// Will not be called
func (img *Picture) OnKeyPress(ebiten.Key) {

}

// Will not be called
func (img *Picture) GetSelectedButtonLabel() string {
	return ""
}
