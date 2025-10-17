package types

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
	UNDEFINED
)

var directionName = map[Direction]string{
	UP:        "up",
	DOWN:      "down",
	LEFT:      "left",
	RIGHT:     "right",
	UNDEFINED: "undefined",
}

func (d Direction) String() string {
	return directionName[d]
}
