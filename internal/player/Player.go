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

func (p *Player) Move(d types.Direction) error {
	actualSpeed := p.Speed
	if p.IsRunning {
		actualSpeed *= 2
	}

	switch d {
	case types.UNDEFINED:
	case types.UP:
		// if u ask yourself why i subtracted 10 here, you asked a very good question, but it works so get on with it :)
		if p.Pos.Y >= (p.Size.Y/2)-10 {
			p.Pos.Y -= actualSpeed
		}
	case types.DOWN:
		if p.Pos.Y <= 790-(p.Size.Y) {
			p.Pos.Y += actualSpeed
		}
	case types.LEFT:
		if p.Pos.X >= (p.Size.X/2)-10 {
			p.Pos.X -= actualSpeed
		}
	case types.RIGHT:
		if p.Pos.X <= 790-(p.Size.X) {
			p.Pos.X += actualSpeed
		}
	}
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
