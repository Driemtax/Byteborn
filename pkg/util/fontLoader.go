package util

import (
	"bytes"
	"image/color"
	"log"

	"github.com/Driemtax/Byteborn/internal/config"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Font struct {
	faceSource *text.GoTextFaceSource
	size       int
}

// Initializes the font by loading the ttf file
func InitFont() *Font {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2P_ttf))

	if err != nil {
		log.Fatalf("Error importing ttf Font %v", err)
	}

	return &Font{
		faceSource: s,
		size:       config.PRESS_START_2P_SIZE,
	}
}

func (font *Font) Draw(screen *ebiten.Image, posX int, posY int, content string, justificationX, justificationY text.Align) {
	// Places text in the middle of the op.GeoM point, in this case the middle of the screen
	x := float64(posX)
	y := float64(posY)
	msg := content
	op := &text.DrawOptions{}
	op.PrimaryAlign = justificationX
	op.SecondaryAlign = justificationY
	op.GeoM.Translate(float64(x), float64(y))
	op.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen, msg, &text.GoTextFace{
		Source: font.faceSource,
		Size:   float64(font.size),
	}, op)
	// ---

}
