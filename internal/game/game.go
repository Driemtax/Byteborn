package game

import (
	"log"
	"os"

	"github.com/Driemtax/Byteborn/internal/config"
	"github.com/Driemtax/Byteborn/internal/debug"
	"github.com/Driemtax/Byteborn/internal/player"
	"github.com/Driemtax/Byteborn/internal/scene"
	"github.com/Driemtax/Byteborn/internal/world"
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
	player  *player.Player
	input   *input.InputState
	world   *world.World
	overlay *debug.Overlay
}

func NewGame() *Game {
	return &Game{
		player:  player.NewPlayer(),
		world:   world.NewWorld(),
		overlay: debug.NewOverlay(),
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
	dt := 1.0 / float64(ebiten.TPS())

	// Update player animation count
	g.player.UpdateAC()
	g.player.UpdateLookDirection()

	g.input = input.GetInputState()

	// Check for debug overlay
	if g.input.TOGGLE_DEBUG {
		g.overlay.Toggle()
	}

	// Check for ESC
	if g.input.ESC {
		os.Exit(0)
	}

	dir, err := g.HandleInput()

	if err != nil {
		log.Fatal(err)
	}
	g.player.Move(dir)

	// Reset running every frame
	g.player.IsRunning = false

	g.world.UpdateCamera(g.player.Pos, dt)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.world.Draw(screen)
	g.player.Draw(screen, g.world.GetCameraPos())
	g.drawDebugOverlay(screen)
}

// drawDebugOverlay feeds the current game state into the overlay and renders it.
// Everything in here is skipped while the overlay is hidden, so it costs nothing
// during normal play.
func (g *Game) drawDebugOverlay(screen *ebiten.Image) {
	if !g.overlay.Visible() {
		return
	}

	o := g.overlay
	o.Reset()

	o.Section("PERFORMANCE")
	o.Float("FPS", ebiten.ActualFPS(), 1)
	o.Float("TPS", ebiten.ActualTPS(), 1)
	o.Int64("Tick", ebiten.Tick())

	o.Section("PLAYER")
	o.Vec2("Pos", g.player.Pos)
	o.Vec2("LastDir", g.player.LastDirection)
	o.Text("Facing", g.player.LookDir().String())
	o.Int("Frame", g.player.AnimationFrame())
	o.Bool("Moving", g.player.IsMoving)
	o.Bool("Running", g.player.IsRunning)

	o.Section("INPUT")
	if g.input != nil {
		o.Text("Keys", g.input.String())
	}

	o.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WIDHT, HEIGHT
}

var _ scene.Scene = (*Game)(nil)
