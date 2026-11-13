package ui

import (
	"github.com/Driemtax/Byteborn/pkg/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Text struct {
	Box
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

func (t *Text) Draw(screen *ebiten.Image) {
	//midX := (t.bounds.Min.X + t.bounds.Max.X) / 2
	//midY := (t.bounds.Min.Y + t.bounds.Max.Y) / 2
	t.font.Draw(screen, t.bounds.Min.X, t.bounds.Min.Y, t.content, t.justificationX, t.justificationY)
}
