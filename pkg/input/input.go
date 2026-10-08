package input

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type InputState struct {
	UP     bool
	DOWN   bool
	RIGHT  bool
	LEFT   bool
	LSHIFT bool
	ESC    bool
}

func GetInputState() *InputState {
	return &InputState{
		UP:     ebiten.IsKeyPressed(ebiten.KeyW),
		DOWN:   ebiten.IsKeyPressed(ebiten.KeyS),
		LEFT:   ebiten.IsKeyPressed(ebiten.KeyA),
		RIGHT:  ebiten.IsKeyPressed(ebiten.KeyD),
		LSHIFT: ebiten.IsKeyPressed(ebiten.KeyShiftLeft),
		ESC:    ebiten.IsKeyPressed(ebiten.KeyEscape),
	}
}
