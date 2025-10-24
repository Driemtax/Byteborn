package types

import "fmt"

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

func (v Vec2) String() string {
	return fmt.Sprintf("Vec2{X: %f, Y: %f}", v.X, v.Y)
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

func (v Vec2) IsHorizontalBetween(left, right Vec2) bool {
	if left.X < v.X && v.X < right.X {
		return true
	}
	return false
}

func (v Vec2) IsVerticalBetween(top, bottom Vec2) bool {
	if top.Y < v.Y && v.Y < bottom.Y {
		return true
	}
	return false
}
