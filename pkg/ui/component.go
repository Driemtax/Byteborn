package ui

import "github.com/hajimehoshi/ebiten/v2"

type Component interface {
	Draw(screen *ebiten.Image, posX int, posY int, width int, height int)
	OnKeyPress(ebiten.Key)
	GetSelectedButtonLabel() string
}
