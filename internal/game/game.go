package game

import (
	"log"
	"time"

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
)

func init() {
	ebiten.SetWindowSize(SCALED_WIDTH, SCALED_HEIGHT)
	ebiten.SetWindowTitle("Byteborn by Archaide")
	ebiten.SetTPS(60)
}

type Game struct {
	player *player.Player
	input  *input.InputState

	// Delta Time Handling
	dt         float64
	dtDuration time.Duration
}

func NewGame() *Game {
	return &Game{
		player: player.NewPlayer(),
		dt:     0.0,
	}
}

func (g *Game) HandleInput() (types.Vec2, error) {
	var err error

	if g.input.LSHIFT {
		g.player.IsRunning = true
	}

	direction := types.NewVector2D(0, 0)

	// Directions: explanation
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

	return direction, err
}

// Updates the delta time every tick.
func (g *Game) updateDT() {
	g.dt = 1.0 / ebiten.ActualTPS()
	g.dtDuration = time.Second / time.Duration(ebiten.ActualTPS())

	// Updates the player sum of delta times
	g.player.DeltaSum += g.dtDuration
}

func (g *Game) Update() error {
	g.updateDT()
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
