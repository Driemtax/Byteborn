package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type InputState struct {
	UP     bool
	DOWN   bool
	RIGHT  bool
	LEFT   bool
	LSHIFT bool
	ESC    bool

	// DEBUG TOGGLE
	TOGGLE_DEBUG bool
}

func GetInputState() *InputState {
	return &InputState{
		UP:     ebiten.IsKeyPressed(ebiten.KeyW),
		DOWN:   ebiten.IsKeyPressed(ebiten.KeyS),
		LEFT:   ebiten.IsKeyPressed(ebiten.KeyA),
		RIGHT:  ebiten.IsKeyPressed(ebiten.KeyD),
		LSHIFT: ebiten.IsKeyPressed(ebiten.KeyShiftLeft),

		ESC:          inpututil.IsKeyJustPressed(ebiten.KeyEscape),
		TOGGLE_DEBUG: inpututil.IsKeyJustPressed(ebiten.KeyF3),
	}
}

// String returns a compact view of the movement keys currently held down, for
// example "W--D^" while running towards the upper right. Each slot is fixed so
// the debug overlay does not jitter in width.
func (i *InputState) String() string {
	keys := [5]byte{'-', '-', '-', '-', '-'}
	if i.UP {
		keys[0] = 'W'
	}
	if i.LEFT {
		keys[1] = 'A'
	}
	if i.DOWN {
		keys[2] = 'S'
	}
	if i.RIGHT {
		keys[3] = 'D'
	}
	if i.LSHIFT {
		keys[4] = '^'
	}
	return string(keys[:])
}
