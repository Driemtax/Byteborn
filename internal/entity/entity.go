package entity

import (
	"image"

	"github.com/Driemtax/Byteborn/internal/config"
	"github.com/Driemtax/Byteborn/pkg/types"
	"github.com/Driemtax/Byteborn/pkg/util"
	"github.com/hajimehoshi/ebiten/v2"
)

type Base struct {
	Size      types.Vector2D
	Pos       types.Vector2D
	Speed     float64
	IsRunning bool

	// asset management
	SpriteSheet    *ebiten.Image
	LookDirection  int
	AnimationCount int
	IsMoving       bool
	LastDirection  types.Vector2D
}

type LookDirection int

const (
	LEFT LookDirection = iota
	DOWN
	RIGHT
	UP
)

func (d LookDirection) String() string {
	switch d {
	case LEFT:
		return "LEFT"
	case DOWN:
		return "DOWN"
	case RIGHT:
		return "RIGHT"
	case UP:
		return "UP"
	}
	return "UNKNOWN"
}

// LookDir returns the direction the entity is currently facing.
// This is the row index into the sprite sheet.
func (e *Base) LookDir() LookDirection {
	return LookDirection(e.LookDirection)
}

// AnimationFrame returns the current column index into the sprite sheet.
// 0 is the standing pose, 1 and 2 are the walking poses.
func (e *Base) AnimationFrame() int {
	return e.AnimationCount
}

const (
	HEIGHT = config.WINDOW_HEIGHT
	WIDTH  = config.WINDOW_WIDTH
	SPEED  = config.ENTITY_SPEED

	TICK_UPDATE = config.TICK_UPDATE
)

func NewEntity() Base {
	return Base{
		Size:           types.NewVector2D(32, 32),
		Pos:            types.NewVector2D(WIDTH/2, HEIGHT/2),
		Speed:          SPEED,
		IsRunning:      false,
		SpriteSheet:    util.LoadAsset("assets/Poke3.png"),
		LookDirection:  int(DOWN),
		AnimationCount: 0,
		IsMoving:       false,
	}
}

// Updates the animationCount based on the tick count and the config on how fast the animation should be played
func (e *Base) UpdateAC() {
	if e.IsMoving {
		if ebiten.Tick()%int64(TICK_UPDATE) == 0 {
			e.AnimationCount = 1 + (e.AnimationCount & 1) // = e.animationCount % 2
		}
	} else {
		e.AnimationCount = 0 // standing position sprite
	}
}

// Updates the lookDirection of the entity based on the last move it made.
// Diagonal movement results in a look direction of left or right
func (e *Base) UpdateLookDirection() {
	switch {
	case e.LastDirection.X == -1:
		e.LookDirection = int(LEFT)
	case e.LastDirection.X == 1:
		e.LookDirection = int(RIGHT)
	case e.LastDirection.Y == -1:
		e.LookDirection = int(UP)
	case e.LastDirection.Y == 1:
		e.LookDirection = int(DOWN)
	}
}

func (e *Base) Move(dir types.Vector2D) error {
	actualSpeed := e.Speed
	if e.IsRunning {
		actualSpeed *= 1.5
	}

	// Normalize the direction. That garantuees, that diagonal movement is NOT faster then horizontal or vertical movement.
	// See docs for further details.
	dir = dir.Normalize()

	// calculate velocity of movement
	v := dir.Mul(actualSpeed)

	newPos := types.NewVector2D(0, 0)

	// calculate the move. Later on there will be a check against the world api if the move is valid.
	newPos.X = e.Pos.X + v.X
	newPos.Y = e.Pos.Y + v.Y

	// apply the move if valid. For now we asume it is
	e.Pos.X = newPos.X
	e.Pos.Y = newPos.Y

	return nil
}

func (e *Base) Draw(screen *ebiten.Image, cameraPos types.Vector2D) {
	// Get the coordinates of the correct sprite from the spriteSheet
	xStart := e.AnimationCount * int(e.Size.X)
	yStart := e.LookDirection * int(e.Size.Y)
	xEnd := xStart + int(e.Size.X)
	yEnd := yStart + int(e.Size.Y)

	// create a rectangle to cut out the sprite of the spriteSheet
	portion := image.Rect(xStart, yStart, xEnd, yEnd)
	pixels := e.SpriteSheet.SubImage(portion).(*ebiten.Image)

	// lastly just draw the image to the screen at the right position
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(e.Pos.X-cameraPos.X, e.Pos.Y-cameraPos.Y)
	screen.DrawImage(pixels, op)
}
