package world

import (
	"github.com/Driemtax/Byteborn/internal/config"
	"github.com/Driemtax/Byteborn/pkg/types"
	"github.com/hajimehoshi/ebiten/v2"
)

// Chunk: A chunk is a piece of the world, it later could contain the map, the enemies and npc's
// that are currently inside of this chunk.
// TODO: Currently it is just a tile map but later it will also contain all the other things...

// Using a single line with tile = y * w + x (w => width)
type Chunk [config.CHUNK_SIZE * config.CHUNK_SIZE]Tile

// Returns a new chunk
// Each tile will reset to the basic tile type which is NONE
func NewChunk() *Chunk {
	newChunk := Chunk{}
	for n := range config.CHUNK_SIZE * config.CHUNK_SIZE {
		newChunk[n] = NewTile()
	}
	return &newChunk
}

// The string method won't be implemented because the default methods does already everything we need

// A single tile is 8bit
// A chunk consists of 16*16=>256 Tiles => 256 * 8 => 256 byte
// A cache line is normaly 64 bytes long
// So i need 4 cache lines for a whole chunk thats not to bad
func (c *Chunk) GetTile(x, y int) Tile {
	// Check if it is out of bounds
	if y*config.CHUNK_SIZE+x < config.CHUNK_SIZE*config.CHUNK_SIZE {
		return c[y+config.CHUNK_SIZE+x]
	}
	// Its out of bounds so we return none because theres nothing...
	return NONE
}

// Print a 2D matrix representation of the chunk
func (c *Chunk) Display() string {
	display := ""
	for i := range config.CHUNK_SIZE * config.CHUNK_SIZE {
		if i%config.CHUNK_SIZE == 0 {
			if i > 0 {
				display += "|\n"
			}
			for range config.CHUNK_SIZE {
				display += "+--"
			}
			display += "+\n"
		}
		display += "|" + c[i].Display()
	}
	display += "|\n"
	for range config.CHUNK_SIZE {
		display += "+--"
	}
	display += "+\n"
	return display
}

func (c *Chunk) Draw(screen *ebiten.Image, tl types.Vec2) {
	for y := range config.CHUNK_SIZE {
		for x := range config.CHUNK_SIZE {
			c.GetTile(x, y).Draw(screen, tl.X+float64(x*config.TILE_SIZE), tl.Y+float64(y*config.TILE_SIZE))
		}
	}
}
