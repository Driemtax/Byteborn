package tile

import (
	"fmt"
	"image/color"

	"github.com/Driemtax/Byteborn/internal/config"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Tile uint8

const (
	NONE Tile = iota
	GRASS
	PATH
	WALL
	Test
	Test2
	Test3
	MAX_TILE_AMOUNT // Has always to be the last enum
)

var TILE_DETAILS = [MAX_TILE_AMOUNT]TileDetail{
	{name: "None", abbreviation: "##", color: color.RGBA{R: 68, G: 1, B: 84, A: 255}},
	{name: "Grass", abbreviation: "GR", color: color.RGBA{R: 65, G: 68, B: 135, A: 255}},
	{name: "Path", abbreviation: "PA", color: color.RGBA{R: 42, G: 120, B: 142, A: 255}},
	{name: "Wall", abbreviation: "WA", color: color.RGBA{R: 34, G: 168, B: 132, A: 255}},
	{name: "Test", abbreviation: "WA", color: color.RGBA{R: 122, G: 209, B: 81, A: 255}},
	{name: "Test2", abbreviation: "WA", color: color.RGBA{R: 189, G: 223, B: 38, A: 255}},
	{name: "Test3", abbreviation: "WA", color: color.RGBA{R: 253, G: 231, B: 37, A: 255}},
}

func NewTile(noiseValue float32) Tile {
	if noiseValue < -0.5 {
		return GRASS
	} else if noiseValue < 0 {
		return PATH
	} else if noiseValue < 0.33 {
		return Test
	} else if noiseValue < 0.45 {
		return Test2
	} else if noiseValue < 0.8 {
		return Test3
	}
	return WALL
}

func (t Tile) String() string {
	return fmt.Sprintf("Tile%+v", TILE_DETAILS[t])
}

func (t Tile) Display() string {
	return TILE_DETAILS[t].Abbreviation()
}

func (t Tile) Draw(screen *ebiten.Image, x, y float64) {
	vector.FillRect(
		screen,
		float32(x),
		float32(y),
		config.TILE_SIZE,
		config.TILE_SIZE,
		TILE_DETAILS[t].Color(),
		true,
	)
}
