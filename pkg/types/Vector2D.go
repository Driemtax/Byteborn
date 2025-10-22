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
