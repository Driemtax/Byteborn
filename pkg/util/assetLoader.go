package util

import (
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func LoadAsset(path string) *ebiten.Image {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening the file..")
		log.Fatal(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		fmt.Println("Error decoding the image..")
		log.Fatal(err)
	}

	// Ebiten.Image is optimized for GPU Usage
	ebitenImg := ebiten.NewImageFromImage(img)
	return ebitenImg
}
