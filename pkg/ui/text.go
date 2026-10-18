package ui

import (
	"github.com/Driemtax/Byteborn/pkg/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Text struct {
	font    *util.Font
	content string

	justificationX text.Align
	justificationY text.Align
}

func (t *Text) Draw(screen *ebiten.Image, posX int, posY int) {
	t.font.Draw(screen, posX, posY, t.content, t.justificationX, t.justificationY)
}
