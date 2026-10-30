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

func NewText(msg string, justificationX text.Align, justificationY text.Align) *Text {
	return &Text{
		font:           util.InitFont(),
		content:        msg,
		justificationX: justificationX,
		justificationY: justificationY,
	}
}

func (t *Text) Draw(screen *ebiten.Image, posX int, posY int, width int, height int) {
	midX := posX + (width / 2)
	midY := posY + (height / 2)
	t.font.Draw(screen, midX, midY, t.content, t.justificationX, t.justificationY)
}

func (t *Text) OnKeyPress(ebiten.Key) {

}

func (t *Text) GetSelectedButtonLabel() string {
	return ""
}
