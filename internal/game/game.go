package game

import (
	"log"

	"github.com/Driemtax/Byteborn/internal/player"
	"github.com/Driemtax/Byteborn/internal/scene"
	"github.com/Driemtax/Byteborn/pkg/types"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func init() {
	ebiten.SetWindowSize(800, 800)
	ebiten.SetWindowTitle("Byteborn by Archaide")
	ebiten.SetTPS(60)
}

type Game struct {
	player *player.Player
	keys   []ebiten.Key
}

func NewGame() *Game {
	return &Game{
		player: player.NewPlayer(),
	}
}

func (g *Game) HandleInput(p *player.Player) error {
	var err error
	// Idea: Instead of iterating i could identify all unique keys in the array and if there is more then one the direction
	// is diagonal. Then i can normalize the movement speed in move()
	for _, k := range g.keys {
		// doesnt work. i could use vectors here too and just multiply the direction by 2 or something
		if k == ebiten.KeyShift {
			g.player.IsRunning = true
		}
		var direction types.Direction
		switch k {
		case ebiten.KeyW:
			direction = types.UP
		case ebiten.KeyA:
			direction = types.LEFT
		case ebiten.KeyS:
			direction = types.DOWN
		case ebiten.KeyD:
			direction = types.RIGHT
		default:
			direction = types.UNDEFINED
		}

		err = p.Move(direction)
		if err != nil {
			log.Fatal(err)
		}
	}

	// reset running
	g.player.IsRunning = false

	return err
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	g.HandleInput(g.player)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 800
}

var _ scene.Scene = (*Game)(nil)
