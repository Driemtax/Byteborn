package game

import (
	"log"

	"github.com/Driemtax/Byteborn/internal/config"
	"github.com/Driemtax/Byteborn/internal/player"
	"github.com/Driemtax/Byteborn/internal/scene"
	"github.com/Driemtax/Byteborn/pkg/input"
	"github.com/Driemtax/Byteborn/pkg/types"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	WIDHT         = config.WINDOW_WIDTH
	HEIGHT        = config.WINDOW_HEIGHT
	SCALE         = config.WINDOW_SCALE
	SCALED_WIDTH  = WIDHT * SCALE
	SCALED_HEIGHT = HEIGHT * SCALE

	TPS = config.TPS
)

func init() {
	ebiten.SetWindowSize(SCALED_WIDTH, SCALED_HEIGHT)
	ebiten.SetWindowTitle("Byteborn by Archaide")
	ebiten.SetTPS(TPS)
}

type Game struct {
	player *player.Player
	input  *input.InputState
}

func NewGame() *Game {
	return &Game{
		player: player.NewPlayer(),
	}
}

func (g *Game) HandleInput() (types.Vec2, error) {
	// reset walking status of player
	g.player.IsMoving = false
	var err error

	if g.input.LSHIFT {
		g.player.IsRunning = true
	}

	direction := types.NewVector2D(0, 0)

	// Directions: TODO: explanation
	if g.input.UP {
		direction = direction.Add(types.NewVector2D(0, -1))
	}

	if g.input.DOWN {
		direction = direction.Add(types.NewVector2D(0, 1))
	}

	if g.input.RIGHT {
		direction = direction.Add(types.NewVector2D(1, 0))
	}

	if g.input.LEFT {
		direction = direction.Add(types.NewVector2D(-1, 0))
	}

	// Set the player IsMoving for walking animation
	if direction.LengthSq() > 0 {
		g.player.IsMoving = true
		g.player.LastDirection = direction
	}

	return direction, err
}

func (g *Game) Update() error {
	// Update player animation count
	g.player.UpdateAC()
	g.player.UpdateLookDirection()

	g.input = input.GetInputState()
	dir, err := g.HandleInput()

	if err != nil {
		log.Fatal(err)
	}
	g.player.Move(dir)

	// Reset running every frame
	g.player.IsRunning = false

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WIDHT, HEIGHT
}

var _ scene.Scene = (*Game)(nil)
