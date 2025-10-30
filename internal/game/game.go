package game

import (
	"log"

	"github.com/Driemtax/Byteborn/internal/input"
	"github.com/Driemtax/Byteborn/internal/player"
	"github.com/Driemtax/Byteborn/internal/scene"
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

func (g *Game) HandleInput() error {
	var err error

	if g.input.LShift {
		g.player.IsRunning = true
	}

	direction := types.NewVector2D(0, 0)
	if g.input.Up {
		direction = direction.Add(types.NewVector2D(0, -1))
	}
	if g.input.Right {
		direction = direction.Add(types.NewVector2D(1, 0))
	}
	if g.input.Down {
		direction = direction.Add(types.NewVector2D(0, 1))
	}
	if g.input.Left {
		direction = direction.Add(types.NewVector2D(-1, 0))
	}

	err = g.player.Move(direction)
	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (g *Game) Update() error {
	g.input = input.GetInputState()
	g.HandleInput()

	// reset running since a player should only be running as long as shift is pressed
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
