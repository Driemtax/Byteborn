package player

import (
	"github.com/Driemtax/Byteborn/internal/entity"
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
