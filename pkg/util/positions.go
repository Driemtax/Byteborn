package util

import (
	"github.com/Driemtax/Byteborn/pkg/types"
)

func IsInBounds(pos types.Vector2D, minX, minY, maxX, maxY float64) bool {
	if pos.X >= minX && pos.X <= maxX && pos.Y >= minY && pos.Y <= maxY {
		return true
	}
	return false
}

func Clamp(pos types.Vector2D, minX, minY, maxX, maxY float64) types.Vector2D {
	if pos.X < minX {
		pos.X = minX
	} else if pos.X > maxX {
		pos.X = maxX
	}

	if pos.Y < minY {
		pos.Y = minY
	} else if pos.Y > maxY {
		pos.Y = maxY
	}
	return pos
}

func IsInPointsSquare(pos types.Vec2, topLeft types.Vec2, botRight types.Vec2) bool {
	if pos.Le(topLeft) && pos.Se(botRight) {
		return true
	}
	return false
}

// SqareCollides: Checks if the sqares a and b overlap or not
// The Idea is to check if the sqares not collide
// There are 4 checks that we need to do and if
// one of them is true the squares do not overlap:
// 1. To far to the right
// 2. To far to the left
// 3. To far on Top
// 4. To far Below
func SquareCollides(atl, abr, btl, bbr types.Vec2) bool {
	if abr.X < btl.X || atl.X > btl.X || atl.Y > btl.Y || abr.Y < btl.Y {
		return false
	}
	return true
}
