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
	"github.com/Driemtax/Byteborn/pkg/ui"
	"github.com/Driemtax/Byteborn/pkg/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	WIDHT         = config.WINDOW_WIDTH
	HEIGHT        = config.WINDOW_HEIGHT
	SCALE         = config.WINDOW_SCALE
	SCALED_WIDTH  = WIDHT * SCALE
	SCALED_HEIGHT = HEIGHT * SCALE

	TPS = config.TPS

	FONT_SIZE = config.PRESS_START_2P_SIZE
	PADDING   = 20
)

type SCENE_STATE int

const (
	MENU SCENE_STATE = iota
	INGAME
)

func init() {
	ebiten.SetWindowSize(SCALED_WIDTH, SCALED_HEIGHT)
	ebiten.SetWindowTitle("Byteborn by Archaide")
	ebiten.SetTPS(TPS)
}

type Game struct {
	font    *util.Font
	player  *player.Player
	input   *input.InputState
	world   *world.World
	overlay *debug.Overlay

	// ui components
	scene_state SCENE_STATE
	components  []ui.Div
}

func NewGame() *Game {
	return &Game{
		font:    util.InitFont(),
		player:  player.NewPlayer(),
		world:   world.NewWorld(),
		overlay: debug.NewOverlay(),

		// UI
		scene_state: INGAME,
		components:  []ui.Div{},
	}
}

func (g *Game) InitGame() {
	g.InitUI()
	g.scene_state = MENU
}

func (g *Game) InitUI() {
	// --- MAIN DIV ---
	mainDiv := ui.NewDiv(config.WINDOW_WIDTH, config.WINDOW_HEIGHT, 0, 0)
	g.components = append(g.components, *mainDiv)
	// ------------------------------------------------------------------------

	// --- TITLE DIV ---
	titleWidth := (config.WINDOW_WIDTH/3)*2 - 2*PADDING
	titleHeight := config.WINDOW_HEIGHT / 3
	titleX := (config.WINDOW_WIDTH / 3) + PADDING
	titleY := 20
	titleDiv := ui.NewDiv(titleWidth, titleHeight, titleX, titleY)
	titleText := ui.NewText("Byteborn", text.AlignCenter, text.AlignCenter)
	titleDiv.AddChild(titleText)
	g.components = append(g.components, *titleDiv)
	// -------------------------------------------------------------------------

	// --- LOGO DIV ---
	logoWidth := (config.WINDOW_WIDTH/3)*1 - PADDING
	logoHeight := config.WINDOW_HEIGHT / 3
	logoX := 20
	logoY := 20
	logoDiv := ui.NewDiv(logoWidth, logoHeight, logoX, logoY)
	g.components = append(g.components, *logoDiv)

	// --- BUTTONS DIV --
	selectionWidth := config.WINDOW_WIDTH / 3
	selectionHeight := (config.WINDOW_HEIGHT/3)*2 - 3*PADDING
	selectionX := (config.WINDOW_WIDTH / 2) - (selectionWidth / 2)
	selectionY := titleY + titleHeight + PADDING
	selectionDiv := ui.NewDiv(selectionWidth, selectionHeight, selectionX, selectionY)
	g.components = append(g.components, *selectionDiv)

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

	// Handle input accordingly to selected scene
	switch g.scene_state {
	case MENU:
		g.handleMenuInput()
	case INGAME:
		g.handleIngameInput(dt)
	}

	return nil
}

func (g *Game) handleMenuInput() {
	// TODO Handle Menu Input, so div.OnKeyPress or something or probably Menu struct which holds the div
	if g.input.ESC {
		os.Exit(0)
	}

	if g.input.ENTER {
		g.scene_state = INGAME
	}
}

func (g *Game) handleIngameInput(dt float64) {
	// Check for ESC
	if g.input.ESC {
		g.scene_state = MENU
	}

	dir, err := g.HandleInput()

	if err != nil {
		log.Fatal(err)
	}
	g.player.Move(dir)

	// Reset running every frame
	g.player.IsRunning = false

	g.world.UpdateCamera(g.player.Pos, dt)
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw components accordingly to the scene state

	switch g.scene_state {
	case MENU:
		g.drawMenu(screen)
	case INGAME:
		g.drawIngame(screen)
	}
}

func (g *Game) drawMenu(screen *ebiten.Image) {
	// draw all divs to the screen
	for _, d := range g.components {
		d.Draw(screen)
	}
}

func (g *Game) drawIngame(screen *ebiten.Image) {
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
