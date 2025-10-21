package types

type Vector2D struct {
	X float64
	Y float64
}

type Vec2 = Vector2D

// Equals: Checks if the vector is equal to another vector
func (v Vector2D) Eq(otherVector Vector2D) bool {
	if v.X == otherVector.X && v.Y == otherVector.Y {
		return true
	}
	return false
}
