package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/Driemtax/Byteborn/internal/scene"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func init() {
	ebiten.SetWindowSize(400, 400)
	ebiten.SetWindowTitle("Byteborn by Archaide")
	ebiten.SetTPS(60)
}

type TestGame struct{}

func (tg TestGame) Update() error {
	return nil
}

func (tg TestGame) Draw(screen *ebiten.Image) {
	vector.FillRect(
		screen,
		180,
		180,
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
	testGame := TestGame{}
	if err := ebiten.RunGame(testGame); err != nil {
		log.Fatal(err)
	}
}
