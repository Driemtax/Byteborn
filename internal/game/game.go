package game

import (
	"image/color"
	"log"

	"github.com/Driemtax/Byteborn/internal/scene"
	"github.com/Driemtax/Byteborn/pkg/types"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func init() {
	ebiten.SetWindowSize(800, 800)
	ebiten.SetWindowTitle("Byteborn by Archaide")
	ebiten.SetTPS(60)
}

type Game struct {
	player *types.Player
	keys   []ebiten.Key
}

func NewGame() *Game {
	return &Game{
		player: types.NewPlayer(),
	}
}

func (g *Game) HandleInput(p *types.Player) error {
	var err error
	// Idea: Instead of iterating i could identify all unique keys in the array and if there is more then one the direction
	// is diagonal. Then i can normalize the movement speed in move()
	for _, k := range g.keys {
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
	return err
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	g.HandleInput(g.player)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	vector.FillRect(
		screen,
		float32(g.player.Posx),
		float32(g.player.Posy),
		float32(g.player.SizeX),
		float32(g.player.SizeY),
		color.RGBA{R: 255, G: 255, B: 255, A: 255},
		true,
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 800
}

var _ scene.Scene = (*Game)(nil)
