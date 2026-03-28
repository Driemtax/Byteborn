package game

import (
	"log"

	"github.com/Driemtax/Byteborn/internal/player"
	"github.com/Driemtax/Byteborn/internal/scene"
	"github.com/Driemtax/Byteborn/pkg/input"
	"github.com/Driemtax/Byteborn/pkg/types"
	"github.com/hajimehoshi/ebiten/v2"
)

func init() {
	ebiten.SetWindowSize(800, 800)
	ebiten.SetWindowTitle("Byteborn by Archaide")
	ebiten.SetTPS(60)
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

func (g *Game) Update() error {
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
	return 800, 800
}

var _ scene.Scene = (*Game)(nil)
