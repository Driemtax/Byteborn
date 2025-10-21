package util

import "github.com/Driemtax/Byteborn/pkg/types"

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
