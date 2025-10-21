package types

type Vector2D struct {
	X float64
	Y float64
}

type Vec2 = Vector2D

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

// Larger then: Checks if the vector is larger then another vector
func (v Vec2) Lt(otherVector Vec2) bool {
	if v.X > otherVector.X && v.Y > otherVector.Y {
		return true
	}
	return false
}

// Larger equals then: Checks if the vector is large or equals then another vector
func (v Vec2) Le(otherVector Vec2) bool {
	if v.X >= otherVector.X && v.Y >= otherVector.Y {
		return true
	}
	return false
}

// Smaller then: Checks if the vector is smaller then another vector
func (v Vec2) St(otherVector Vec2) bool {
	if v.X < otherVector.X && v.Y < otherVector.Y {
		return true
	}
	return false
}

// Smaller euqals then: Checks if the vector is equals or smaller then another vector
func (v Vec2) Se(otherVector Vec2) bool {
	if v.X <= otherVector.X && v.Y <= otherVector.Y {
		return true
	}
	return false
}
