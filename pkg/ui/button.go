package ui

import (
	"image/color"

	"github.com/Driemtax/Byteborn/pkg/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	BUTTON_PADDING = 10
)

type ButtonManager struct {
	Buttons        []Button
	selectedButton int
}

type Button struct {
	Id         int
	Text       *Text
	IsSelected bool
}

// Create a new button manager with no buttons in it and none selected represented by -1.
func NewButtonManager() *ButtonManager {
	return &ButtonManager{
		Buttons:        []Button{},
		selectedButton: -1,
	}
}

// Adds a new Button to the button manager
func (bm *ButtonManager) AddNewButton(content string, isSelected bool) {
	bm.Buttons = append(bm.Buttons, *NewButton(content, isSelected))

	// if the new button should be selected update the button manager accordingly.
	if isSelected {
		id := len(bm.Buttons) - 1
		bm.selectedButton = id
	}
}

func (bm *ButtonManager) OnKeyPress(key ebiten.Key) {
}

func (bm *ButtonManager) Draw(screen *ebiten.Image, posX int, posY int, width int, height int) {
	buttonCount := len(bm.Buttons)
	bHeight := height / buttonCount

	for i, b := range bm.Buttons {
		// TODO: Calculate position for each button and draw selected icon
		x := posX + BUTTON_PADDING
		y := posY + (bHeight * i)

		b.Draw(screen, x, y, width-2*BUTTON_PADDING, bHeight)
	}
}

func NewButton(content string, isSelected bool) *Button {
	return &Button{
		Text:       &Text{font: util.InitFont(), content: content, justificationX: text.AlignCenter, justificationY: text.AlignCenter},
		IsSelected: isSelected,
	}
}

func (b *Button) OnClick() {

}

func (b *Button) Draw(screen *ebiten.Image, posX int, posY int, width int, height int) {
	vector.StrokeRect(screen, float32(posX), float32(posY), float32(width), float32(height), 1, color.White, false)

	if b.IsSelected {
		// TODO: Draw selected marker at left side of button so at beginning of Button box + 5px or somewhat
	}

	// Label of button
	b.Text.Draw(screen, posX, posY, width, height)
}
