package player

import (
	"image"

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

const (
	HEIGHT = config.WINDOW_HEIGHT
	WIDHT  = config.WINDOW_WIDTH
	SPEED  = config.PLAYER_SPEED

	TICK_UPDATE = config.TICK_UPDATE
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
		IsMoving:       false,
	}
}

// Updates the animationCount based on the tick count and the config on how fast the animation should be played
func (p *Player) UpdateAC() {
	if p.IsMoving {
		if ebiten.Tick()%int64(TICK_UPDATE) == 0 {
			p.animationCount = 1 + (p.animationCount & 1) // = p.animationCount % 2
		}
	} else {
		p.animationCount = 0 // standing position sprite
	}
}

// Updates the lookDirection of the player based on the last move it made.
// Diagonal movement results in a look direction of left or right
func (p *Player) UpdateLookDirection() {
	switch {
	case p.LastDirection.X == -1:
		p.lookDirection = int(LEFT)
	case p.LastDirection.X == 1:
		p.lookDirection = int(RIGHT)
	case p.LastDirection.Y == -1:
		p.lookDirection = int(UP)
	case p.LastDirection.Y == 1:
		p.lookDirection = int(DOWN)
	}
}

func (p *Player) Move(dir types.Vector2D) error {
	actualSpeed := p.Speed
	if p.IsRunning {
		actualSpeed *= 1.5
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
