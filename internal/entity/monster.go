package entity

import (
	"github.com/Driemtax/Byteborn/pkg/types"
	"github.com/Driemtax/Byteborn/pkg/util"
)

type Monster struct {
	Base
	Name string
}

func NewMonster(name string) *Monster {
	pen := &Monster{
		Base: NewEntity(),
		Name: name,
	}

	pen.SpriteSheet = util.LoadAsset("assets/penguin.png")
	pen.Pos = types.NewVector2D((WIDTH/2)-128, HEIGHT/2)
	pen.LookDirection = int(RIGHT)

	return pen
}

func (m *Monster) Update(targetPos types.Vector2D) {
	// First we update the animation count and look direction based on last move.
	m.UpdateAC()
	m.UpdateLookDirection()

	m.IsMoving = false

	// Simplest idea
	difference := targetPos.Sub(m.Pos)
	goalDifferenceToPlayer := 64.0
	absoluteDistanceToPlayer := difference.Absolute()

	dir := types.NewVector2D(0, 0)
	// if x is negative we go to the left
	if difference.X < 0 && absoluteDistanceToPlayer.X >= goalDifferenceToPlayer {
		dir.X = -1
	} else if difference.X > 0 && absoluteDistanceToPlayer.X >= goalDifferenceToPlayer {
		dir.X = 1
	}

	// same for y
	if difference.Y < 0 && absoluteDistanceToPlayer.Y >= goalDifferenceToPlayer {
		dir.Y = -1
	} else if difference.Y > 0 && absoluteDistanceToPlayer.Y >= goalDifferenceToPlayer {
		dir.Y = 1
	}

	if dir.LengthSq() > 0 {
		m.IsMoving = true
		m.LastDirection = dir
	}

	m.Move(dir)
}
