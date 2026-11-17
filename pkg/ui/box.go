package ui

import "image"

type Box struct {
	bounds image.Rectangle
}

func (b *Box) Layout(r image.Rectangle) {
	b.bounds = r
}

func (b *Box) Bounds() image.Rectangle {
	return b.bounds
}
