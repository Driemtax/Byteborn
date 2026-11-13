package ui

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Component interface {
	Layout(bounds image.Rectangle)
	Draw(screen *ebiten.Image)
	Bounds() image.Rectangle
}

// This checks at compile time that all components actually implement the interface correctly. Bounds is implemented by box.go
// and composition inherits it to all used components
var (
	_ Component = (*Text)(nil)
	_ Component = (*Picture)(nil)
	_ Component = (*Triangle)(nil)
	_ Component = (*Button)(nil)
	_ Component = (*Div)(nil)
)

type KeyPressHandler interface {
	OnKeyPress(key ebiten.Key) bool
}
