package input

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type InputState struct {
	Up     bool
	Right  bool
	Down   bool
	Left   bool
	LShift bool
	Esc    bool
}

func GetInputState() *InputState {
	return &InputState{
		Up:     ebiten.IsKeyPressed(ebiten.KeyW),
		Right:  ebiten.IsKeyPressed(ebiten.KeyD),
		Down:   ebiten.IsKeyPressed(ebiten.KeyS),
		Left:   ebiten.IsKeyPressed(ebiten.KeyA),
		LShift: ebiten.IsKeyPressed(ebiten.KeyShiftLeft),
		Esc:    ebiten.IsKeyPressed(ebiten.KeyEscape),
	}
}
