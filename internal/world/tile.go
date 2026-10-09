package world

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
	{name: "None", abbreviation: "##", color: color.RGBA{R: 0, G: 0, B: 0, A: 255}},
	{name: "Grass", abbreviation: "GR", color: color.RGBA{R: 0, G: 230, B: 0, A: 255}},
	{name: "Path", abbreviation: "PA", color: color.RGBA{R: 150, G: 150, B: 10, A: 255}},
	{name: "Wall", abbreviation: "WA", color: color.RGBA{R: 255, G: 0, B: 0, A: 255}},
	{name: "Test", abbreviation: "WA", color: color.RGBA{R: 255, G: 255, B: 0, A: 255}},
	{name: "Test2", abbreviation: "WA", color: color.RGBA{R: 0, G: 255, B: 255, A: 255}},
	{name: "Test3", abbreviation: "WA", color: color.RGBA{R: 20, G: 20, B: 20, A: 255}},
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
