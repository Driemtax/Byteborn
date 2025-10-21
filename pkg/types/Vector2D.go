package types

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

// Equals: Checks if the vector is equal to another vector
func (v Vector2D) Eq(otherVector Vector2D) bool {
	if v.X == otherVector.X && v.Y == otherVector.Y {
		return true
	}
	return false
}
