package game

import (
	"os"

	"github.com/Driemtax/Byteborn/internal/config"
	"github.com/Driemtax/Byteborn/internal/debug"
	"github.com/Driemtax/Byteborn/internal/entity"
	"github.com/Driemtax/Byteborn/internal/player"
	"github.com/Driemtax/Byteborn/internal/scene"
	"github.com/Driemtax/Byteborn/internal/world"
	"github.com/Driemtax/Byteborn/pkg/input"
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
	player     *player.Player
	testMonser *entity.Monster
	input      *input.InputState
	world      *world.World
	overlay    *debug.Overlay
}

func NewGame() *Game {
	return &Game{
		player:     player.NewPlayer(),
		testMonser: entity.NewMonster("penguin"),
		world:      world.NewWorld(),
		overlay:    debug.NewOverlay(),
	}
}

func (g *Game) Update() error {
	dt := 1.0 / float64(ebiten.TPS())

	g.input = input.GetInputState()

	// Check for debug overlay
	if g.input.TOGGLE_DEBUG {
		g.overlay.Toggle()
	}

	// Check for ESC
	if g.input.ESC {
		os.Exit(0)
	}

	g.player.Update(*g.input)
	g.testMonser.Update(g.player.Pos)

	g.world.UpdateCamera(g.player.Pos, dt)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.world.Draw(screen)
	g.player.Draw(screen, g.world.GetCameraPos())
	g.testMonser.Draw(screen, g.world.GetCameraPos())
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
