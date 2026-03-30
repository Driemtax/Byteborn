package player

import (
	"image"

	"github.com/Driemtax/Byteborn/pkg/types"
	"github.com/Driemtax/Byteborn/pkg/util"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	Size      types.Vector2D
	Pos       types.Vector2D
	Speed     float64
	IsRunning bool

	// asset management
	spriteSheet   *ebiten.Image
	lookDirection int
	frameCount    int
	isMoving      bool
}

type LookDirection int

const (
	LEFT LookDirection = iota
	DOWN
	RIGHT
	UP
)

func NewPlayer() *Player {
	return &Player{
		Size:          types.NewVector2D(32, 32),
		Pos:           types.NewVector2D(380, 380),
		Speed:         10.0,
		IsRunning:     false,
		spriteSheet:   util.LoadAsset("assets/Poke3.png"),
		lookDirection: int(DOWN),
		frameCount:    0,
		isMoving:      false,
	}
}

func (p *Player) Move(dir types.Vector2D) error {
	actualSpeed := p.Speed
	if p.IsRunning {
		actualSpeed *= 2
	}

	// Normalize the direction. That garantuees, that diagonal movement is NOT faster as horizontal or vertical movement.
	// See docs for further details.
	dir = dir.Normalize()

	// calculate velocity of movement
	v := dir.Mul(actualSpeed)

	newPos := types.NewVector2D(0, 0)

	// Apply the move. This function holds for all directions. See documentation for further eplanations on this formula
	newPos.X = max(0, min(p.Pos.X+v.X, 800-p.Size.X)) // TODO: Add variable for Window width
	newPos.Y = max(0, min(p.Pos.Y+v.Y, 800-p.Size.Y)) // TODO: Add variable for Window height

	// TODO: Check if newPos is legit with map api (cant walk on water etc.)
	// For now we assume it is.
	p.Pos.X = newPos.X
	p.Pos.Y = newPos.Y

	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	// Get the coordinates of the correct sprite from spriteSheet
	xStart := p.frameCount * int(p.Size.X)
	yStart := p.lookDirection * int(p.Size.Y)
	xEnd := xStart + int(p.Size.X)
	yEnd := yStart + int(p.Size.Y)

	// create a Rectangle to cut out the sprite of the spriteSheet
	portion := image.Rect(xStart, yStart, xEnd, yEnd)
	pixels := p.spriteSheet.SubImage(portion).(*ebiten.Image)

	// lastly just draw the image to the screen at the right position
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.Pos.X, p.Pos.Y)
	screen.DrawImage(pixels, op)
}
