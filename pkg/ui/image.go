package ui

import (
	"github.com/Driemtax/Byteborn/pkg/util"
	"github.com/hajimehoshi/ebiten/v2"
)

type Picture struct {
	Box
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

func (img *Picture) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	picBounds := img.bounds
	imgBounds := img.file.Bounds()
	widthDiff := picBounds.Dx() - imgBounds.Dx()
	centerX := (widthDiff / 2) - 15
	op.GeoM.Scale(2.0, 2.0)
	op.GeoM.Translate(float64(centerX), 0.0)

	screen.DrawImage(img.file, op)
	// Circle for debugging the centering
	//vector.FillCircle(screen, float32(centerX), float32(midY), 10.0, color.White, false)
}
