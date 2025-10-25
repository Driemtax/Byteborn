package types

import (
	"fmt"
	"math"
)

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

// Add returns the vector sum of v and other (v + other).
// It does not modify the original vector v.
func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{v.X + other.X, v.Y + other.Y}
}

// Sub returns the vector difference of v and other (v - other).
// It does not modify the original vector v.
func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{v.X - other.X, v.Y - other.Y}
}

// Mul returns the vector v scaled by the given scalar factor (v * scalar).
// It does not modify the original vector v.
func (v Vec2) Mul(scalar float64) Vec2 {
	return Vec2{v.X * scalar, v.Y * scalar}
}

// LengthSq returns the squared magnitude (length) of the vector (v.X*v.X + v.Y*v.Y).
// This is computationally cheaper than Len() as it avoids the square root calculation.
// Useful for comparing vector lengths.
func (v Vec2) LengthSq() float64 {
	return v.X*v.X + v.Y*v.Y
}

// Len returns the magnitude (length) of the vector (sqrt(v.X*v.X + v.Y*v.Y)).
// If you only need to compare lengths, use LengthSq() for better performance.
func (v Vec2) Len() float64 {
	// Directly use LengthSq() for clarity and potential minor optimization reuse.
	return math.Sqrt(v.LengthSq())
}

// Dot returns the dot product (scalar product) of vectors v and other.
// The dot product is calculated as v.X*other.X + v.Y*other.Y.
// Useful for calculating angles between vectors or projecting one vector onto another.
func (v Vec2) Dot(other Vec2) float64 {
	return v.X*other.X + v.Y*other.Y
}

func (v Vec2) Rotate(angleRad float64) Vec2 {
	cosA := math.Cos(angleRad)
	sinA := math.Sin(angleRad)
	newX := v.X*cosA - v.Y*sinA
	newY := v.X*sinA + v.Y*cosA
	return Vec2{X: newX, Y: newY}
}

// Normalize returns a unit vector (a vector with length 1) pointing in the same direction as v.
// If the original vector v has a length of 0, it returns a zero vector {0, 0}.
// It does not modify the original vector v.
func (v Vec2) Normalize() Vec2 {
	lenSq := v.LengthSq()
	if lenSq == 0 {
		// Cannot normalize a zero-length vector, return zero vector.
		return Vec2{0, 0}
	}
	len := math.Sqrt(lenSq) // Calculate length only if non-zero
	return Vec2{v.X / len, v.Y / len}
}

// Trunc truncates the values of the vector to their integer value
func (v Vec2) Trunc() Vec2 {
	return Vec2{X: math.Trunc(v.X), Y: math.Trunc(v.Y)}
}
