package types

type Direction int8

const (
	UNDEFINED Direction = iota
	UP
	DOWN
	LEFT
	RIGHT
)

var directionName = map[Direction]string{
	UNDEFINED: "undefined",
	UP:        "up",
	DOWN:      "down",
	LEFT:      "left",
	RIGHT:     "right",
}

func (d Direction) String() string {
	return directionName[d]
}
