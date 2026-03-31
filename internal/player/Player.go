package player

import (
	"image"
	"time"

	"github.com/Driemtax/Byteborn/internal/config"
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
	spriteSheet    *ebiten.Image
	lookDirection  int
	animationCount int
	DeltaSum       time.Duration
	isMoving       bool
}

type LookDirection int

const (
	LEFT LookDirection = iota
	DOWN
	RIGHT
	UP
)

const (
	HEIGHT = config.WINDOW_HEIGHT
	WIDHT  = config.WINDOW_WIDTH
	SPEED  = config.PLAYER_SPEED

	ANIMATION_UPDATE_INTERVALL = config.ANIMATION_UPDATE_INTERVALL
)

func NewPlayer() *Player {
	return &Player{
		Size:           types.NewVector2D(32, 32),
		Pos:            types.NewVector2D(WIDHT/2, HEIGHT/2),
		Speed:          SPEED,
		IsRunning:      false,
		spriteSheet:    util.LoadAsset("assets/Poke3.png"),
		lookDirection:  int(DOWN),
		animationCount: 0,
		DeltaSum:       time.Duration(0),
		isMoving:       false,
	}
}

// Updates the animationCount based on the deltaSum and the config on how fast the animation should be played
func (p *Player) updateAC() {
	// if the last time since the animation was updated is longer then 500ms, then we update the animation once more
	// and reset the sum of delta times to gather 500ms again.
	if p.DeltaSum.Milliseconds() >= ANIMATION_UPDATE_INTERVALL {
		p.animationCount = (p.animationCount + 1) % 3
		p.DeltaSum = 0.0
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
	newPos.X = max(0, min(p.Pos.X+v.X, WIDHT-p.Size.X))  // TODO: Add variable for Window width
	newPos.Y = max(0, min(p.Pos.Y+v.Y, HEIGHT-p.Size.Y)) // TODO: Add variable for Window height

	// TODO: Check if newPos is legit with map api (cant walk on water etc.)
	// For now we assume it is.
	p.Pos.X = newPos.X
	p.Pos.Y = newPos.Y

	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	// Get the coordinates of the correct sprite from spriteSheet
	xStart := p.animationCount * int(p.Size.X)
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
