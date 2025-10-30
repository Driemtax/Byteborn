package types

import (
	"math"
)

type Vector2D struct {
	X float64
	Y float64
}

func NewVector2D(x, y float64) Vector2D {
	return Vector2D{
		X: x,
		Y: y,
	}
}

func (v Vector2D) Add(otherVec Vector2D) Vector2D {
	result := v
	result.X = v.X + otherVec.X
	result.Y = v.Y + otherVec.Y

	return result
}

// Normalizes a 2d vector, which means every value is between 0 and 1 afterwards
func (v Vector2D) Normalize() Vector2D {
	// Get the length of the vector first. Length = |v|
	length := math.Sqrt((v.X * v.X) + (v.Y * v.Y))

	if length == 0 {
		return Vector2D{0, 0}
	}

	return Vector2D{v.X / length, v.Y / length}
}
