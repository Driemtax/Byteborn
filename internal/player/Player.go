package player

import (
	"image/color"

	"github.com/Driemtax/Byteborn/pkg/types"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Player struct {
	Size      types.Vector2D
	Pos       types.Vector2D
	Speed     float64
	IsRunning bool
}

func NewPlayer() *Player {
	return &Player{
		Size:  types.NewVector2D(40, 40),
		Pos:   types.NewVector2D(380, 380),
		Speed: 10.0,
	}
}

func (p *Player) Move(dir types.Vector2D) error {
	actualSpeed := p.Speed
	if p.IsRunning {
		actualSpeed *= 2
	}

	// Normalize the direction. That garantuees, that diagonal movement is NOT faster as horizontal or vertical movement.
	// See docs for further details.
	dir = dir.Normalize()

	// calculate velocity of movement
	v := dir.Mul(actualSpeed)

	newPos := types.NewVector2D(0, 0)

	// Apply the move. This function holds for all directions. See documentation for further eplanations on this formula
	newPos.X = max(0, min(p.Pos.X+v.X, 800-p.Size.X)) // TODO: Add variable for Window width
	newPos.Y = max(0, min(p.Pos.Y+v.Y, 800-p.Size.Y)) // TODO: Add variable for Window height

	// TODO: Check if newPos is legit with map api (cant walk on water etc.)
	// For now we assume it is.
	p.Pos.X = newPos.X
	p.Pos.Y = newPos.Y

	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	vector.FillRect(
		screen,
		float32(p.Pos.X),
		float32(p.Pos.Y),
		float32(p.Size.X),
		float32(p.Size.Y),
		color.RGBA{R: 255, G: 255, B: 255, A: 255},
		true,
	)
}
