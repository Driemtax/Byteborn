package main

import (
	"fmt"
	"log"

	"github.com/Driemtax/Byteborn/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	fmt.Println("Welcome to ByteBorn")
	testGame := game.NewGame()

	testGame.InitGame()
	if err := ebiten.RunGame(testGame); err != nil {
		log.Fatal(err)
	}
}
