package util

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	faceSource *text.GoTextFaceSource
)

// Initializes the font by loading the ttf file
func InitFont() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2P_ttf))

	if err != nil {
		log.Fatalf("Error importing ttf Font %v", err)
	}

	faceSource = s
}

func GetFaceSource() *text.GoTextFaceSource {
	return faceSource
}
