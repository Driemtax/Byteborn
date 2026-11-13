package entity

import (
	"github.com/Driemtax/Byteborn/pkg/types"
	"github.com/Driemtax/Byteborn/pkg/util"
)

type Penguin struct {
	Base
}

func NewPenguin() *Penguin {
	pen := &Penguin{
		Base: NewEntity(),
	}

	pen.SpriteSheet = util.LoadAsset("assets/penguin.png")
	pen.Pos = types.NewVector2D((WIDTH/2)-128, HEIGHT/2)
	pen.LookDirection = int(RIGHT)

	return pen
}
