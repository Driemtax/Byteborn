package player

import (
	"github.com/Driemtax/Byteborn/internal/entity"
	"github.com/Driemtax/Byteborn/pkg/input"
	"github.com/Driemtax/Byteborn/pkg/types"
)

type Player struct {
	entity.Base
}

func NewPlayer() *Player {
	p := &Player{
		Base: entity.NewEntity(),
	}

	return p
}

func (p *Player) Update(input input.InputState) {
	p.UpdateAC()
	p.UpdateLookDirection()

	dir := p.HandleInput(input)

	p.Move(dir)

	p.IsRunning = false
}

func (p *Player) HandleInput(input input.InputState) types.Vector2D {
	// reset walking status of player
	p.IsMoving = false

	if input.LSHIFT {
		p.IsRunning = true
	}

	direction := types.NewVector2D(0, 0)

	if input.UP {
		direction = direction.Add(types.NewVector2D(0, -1))
	}

	if input.DOWN {
		direction = direction.Add(types.NewVector2D(0, 1))

	}

	if input.RIGHT {
		direction = direction.Add(types.NewVector2D(1, 0))

	}

	if input.LEFT {
		direction = direction.Add(types.NewVector2D(-1, 0))
	}

	if direction.LengthSq() > 0 {
		p.IsMoving = true
		p.LastDirection = direction
	}

	return direction
}
