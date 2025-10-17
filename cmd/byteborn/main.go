package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/Driemtax/Byteborn/internal/scene"
	"github.com/Driemtax/Byteborn/pkg/types"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func init() {
	ebiten.SetWindowSize(400, 400)
	ebiten.SetWindowTitle("Byteborn by Archaide")
	ebiten.SetTPS(60)
}

type TestGame struct {
	player *types.Player
	keys   []ebiten.Key
}

func (tg TestGame) Update() error {
	//fmt.Printf("Current Pos: X:%d, Y:%d \n", tg.player.Posx, tg.player.Posy)
	tg.keys = inpututil.AppendPressedKeys(tg.keys[:0])
	tg.HandleInput(tg.player)
	//fmt.Printf("Pos after: X:%d, Y:%d \n", tg.player.Posx, tg.player.Posy)
	return nil
}

func (tg TestGame) HandleInput(p *types.Player) error {
	var err error
	for _, k := range tg.keys {
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
		}

		err = p.Move(direction)
		if err != nil {
			log.Fatal(err)
		}
	}
	return err
}

func (tg TestGame) Draw(screen *ebiten.Image) {
	vector.FillRect(
		screen,
		float32(tg.player.Posx),
		float32(tg.player.Posy),
		40,
		40,
		color.RGBA{R: 255, G: 255, B: 255, A: 255},
		true,
	)
}

func (tg TestGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 400, 400
}

var _ scene.Scene = (*TestGame)(nil)

func main() {
	fmt.Println("Welcome to ByteBorn")
	testGame := TestGame{
		player: types.NewPlayer(),
	}
	if err := ebiten.RunGame(testGame); err != nil {
		log.Fatal(err)
	}
}
