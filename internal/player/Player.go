package player

import (
	"image/color"

	"github.com/Driemtax/Byteborn/pkg/types"
	"github.com/Driemtax/Byteborn/pkg/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Player struct {
	Size      types.Vector2D
	Pos       types.Vector2D
	Speed     float64
	IsRunning bool
}

func NewPlayer() *Player {
	return &Player{
		Size:  types.NewVector2D(40, 40),
		Pos:   types.NewVector2D(380, 380),
		Speed: 10.0,
	}
}

func (p *Player) Update() {

}

// Moves the player based on a 2d direction vector
// UP = 	(0,-1)
// RIGHT = 	(1,0)
// DOWN = 	(0,1)
// LEFT = 	(-1,0)
// UP_RIGHT = UP + RIGHT
// ...
func (p *Player) Move(dir types.Vector2D, world *util.Map) error {
	actualSpeed := p.Speed
	if p.IsRunning {
		actualSpeed *= 2
	}

	var newPos types.Vector2D
	switch dir {
	// UP
	case types.NewVector2D(0, -1):
		if p.checkUp(actualSpeed) {
			newPos.X = p.Pos.X
			newPos.Y = p.Pos.Y + dir.Y*actualSpeed
		}
	// RIGHT
	case types.NewVector2D(1, 0):
		if p.checkRight() {
			newPos.X = p.Pos.X + dir.X*actualSpeed
			newPos.Y = p.Pos.Y
		}
	// DOWN
	case types.NewVector2D(0, 1):
		if p.checkDown() {
			newPos.X = p.Pos.X
			newPos.Y = p.Pos.Y + dir.Y*actualSpeed
		}
	// LEFT
	case types.NewVector2D(-1, 0):
		if p.checkLeft(actualSpeed) {
			newPos.X = p.Pos.X + dir.X*actualSpeed
			newPos.Y = p.Pos.Y
		}
	// UP_RIGHT
	case types.NewVector2D(1, -1):
		dir = dir.Normalize()
		if p.checkRight() && p.checkUp(actualSpeed) {
			newPos.X = p.Pos.X + dir.X*actualSpeed
			newPos.Y = p.Pos.Y + dir.Y*actualSpeed
		}
	// UP_LEFT
	case types.NewVector2D(-1, -1):
		dir = dir.Normalize()
		if p.checkUp(actualSpeed) && p.checkLeft(actualSpeed) {
			newPos.X = p.Pos.X + dir.X*actualSpeed
			newPos.Y = p.Pos.Y + dir.Y*actualSpeed
		}
	// DOWN_RIGHT
	case types.NewVector2D(1, 1):
		dir = dir.Normalize()
		if p.checkDown() && p.checkRight() {
			newPos.X = p.Pos.X + dir.X*actualSpeed
			newPos.Y = p.Pos.Y + dir.Y*actualSpeed
		}
	// DOWN_LEFT
	case types.NewVector2D(-1, 1):
		dir = dir.Normalize()
		if p.checkDown() && p.checkLeft(actualSpeed) {
			newPos.X = p.Pos.X + dir.X*actualSpeed
			newPos.Y = p.Pos.Y + dir.Y*actualSpeed
		}
	}

	if p.checkValidMove(newPos, world) {
		p.Pos.X = newPos.X
		p.Pos.Y = newPos.Y
	}

	return nil
}

// Checks if the tile the player moves is a valid position, e.g. the player can't walk on water
// TODO: currently i move 10 pixel, but i only check the new position, not what is in between
func (p *Player) checkValidMove(newPos types.Vector2D, world *util.Map) bool {
	result := false

	if world.Tiles[int(newPos.Y)][int(newPos.X)] == 1 || world.Tiles[int(newPos.Y)][int(newPos.X)] == 2 {
		result = true
	}

	return result
}

// Checks if there is enough space on the map to move the player upwards based on his position
func (p *Player) checkUp(speed float64) bool {
	return p.Pos.Y >= (p.Size.Y/2)-speed
}

// Checks if there is enough space on the map to move the player right based on his position
func (p *Player) checkRight() bool {
	return p.Pos.X <= 790-(p.Size.X)
}

// Checks if there is enough space on the map to move the player down based on his position
func (p *Player) checkDown() bool {
	return p.Pos.Y <= 790-(p.Size.Y)
}

// Checks if there is enough space on the map to move the player left based on his position
func (p *Player) checkLeft(speed float64) bool {
	return p.Pos.X >= (p.Size.X/2)-speed
}

func (p *Player) Draw(screen *ebiten.Image) {
	vector.FillRect(
		screen,
		float32(p.Pos.X),
		float32(p.Pos.Y),
		float32(p.Size.X),
		float32(p.Size.Y),
		color.RGBA{R: 255, G: 255, B: 255, A: 255},
		true,
	)
}
