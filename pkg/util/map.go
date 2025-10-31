package util

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Map struct {
	Tiles [800][800]int
}

const (
	UNDEFINED int8 = iota
	GRASS
	PATH
	HOUSE
	WATER
)

func NewMap() *Map {
	return &Map{
		Tiles: LoadTiles(),
	}
}

func (m *Map) Draw(screen *ebiten.Image) {
	for i, x := range m.Tiles {
		for j, y := range x {
			var color color.RGBA
			switch y {
			case 1:
				color.R = 80
				color.G = 200
				color.B = 120
			case 2:
				color.R = 220
				color.G = 210
				color.B = 160
			case 3:
				color.R = 200
				color.G = 60
				color.B = 60
			case 4:
				color.R = 70
				color.G = 130
				color.B = 220
			default:
				color.R, color.G, color.B = 128, 128, 128
			}

			vector.FillRect(screen, float32(i), float32(j), 1, 1, color, false)
		}
	}
}

func LoadTiles() [800][800]int {
	result := [800][800]int{}

	// i = row, j = col
	for i := 0; i <= 799; i++ {
		for j := 0; j <= 799; j++ {
			condition := false
			if i >= 360 && i <= 440 {
				result[i][j] = 2
				condition = true
			}
			if j < 100 && i < 100 {
				result[i][j] = 3
				condition = true
			}
			if i > 699 && j > 699 {
				result[i][j] = 3
				condition = true
			}
			if i >= j-2 && i <= j+2 {
				result[i][j] = 4
				condition = true
			}
			if condition == false {
				result[i][j] = 1
			}
		}
	}

	return result

}
