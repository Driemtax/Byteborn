package ui

import (
	"image"

	"github.com/Driemtax/Byteborn/internal/config"
	"github.com/Driemtax/Byteborn/pkg/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	BUTTON_PADDING = 10
)

type ButtonManager struct {
	Box

	Buttons        []Button
	selectedButton int
}

type Button struct {
	Box
	Id         int
	Label      string
	Text       *Text
	IsSelected bool
	selector   *Triangle

	action func()
}

func (bm *ButtonManager) GetSelectedButtonLabel() string {
	return bm.Buttons[bm.selectedButton].GetSelectedButtonLabel()
}

func (b *Button) GetSelectedButtonLabel() string {
	return b.Text.content
}

// Create a new button manager with no buttons in it and none selected represented by -1.
func NewButtonManager() *ButtonManager {
	return &ButtonManager{
		Buttons:        []Button{},
		selectedButton: -1,
	}
}

// Adds a new Button to the button manager
func (bm *ButtonManager) AddNewButton(content string, isSelected bool, action func()) {
	bm.Buttons = append(bm.Buttons, *NewButton(content, isSelected, action))

	// if the new button should be selected update the button manager accordingly.
	if isSelected {
		id := len(bm.Buttons) - 1
		bm.selectedButton = id
	}
}

func (bm *ButtonManager) SelectPrevious() {
	// Check if there is one more button
	if bm.selectedButton >= 1 {
		bm.Buttons[bm.selectedButton].IsSelected = false
		bm.selectedButton -= 1
		bm.Buttons[bm.selectedButton].IsSelected = true
	}
}

func (bm *ButtonManager) SelectNext() {
	// Check if we there is one more button to the top
	if bm.selectedButton < len(bm.Buttons)-1 {
		bm.Buttons[bm.selectedButton].IsSelected = false
		bm.selectedButton += 1
		bm.Buttons[bm.selectedButton].IsSelected = true
	}
}

func (bm *ButtonManager) ActivateSelected() {
	if bm.selectedButton < 0 || bm.selectedButton >= len(bm.Buttons) {
		return
	}

	if a := bm.Buttons[bm.selectedButton].action; a != nil {
		a()
	}
}

// This overrides the Layout Method of the Box composition. The button Manager resizes all buttons equally in its space.
func (bm *ButtonManager) Layout(bounds image.Rectangle) {
	bm.bounds = bounds

	buttonCount := len(bm.Buttons)
	bHeight := bm.bounds.Dy() / buttonCount

	for i, b := range bm.Buttons {
		x := bm.bounds.Min.X + BUTTON_PADDING
		y := bm.bounds.Min.Y + (bHeight * i)

		b.Layout(image.Rect(x, y, x+bm.bounds.Dx()+BUTTON_PADDING, y+bHeight))
	}
}

func (bm *ButtonManager) Draw(screen *ebiten.Image) {
	for _, b := range bm.Buttons {
		b.Draw(screen)
	}
}

func NewButton(content string, isSelected bool, action func()) *Button {
	return &Button{
		Text:       &Text{font: util.InitFont(), content: content, justificationX: text.AlignCenter, justificationY: text.AlignCenter},
		IsSelected: isSelected,
		selector:   NewTriangle(config.PRESS_START_2P_SIZE, config.PRESS_START_2P_SIZE),
		action:     action,
	}
}

func (b *Button) Layout(bounds image.Rectangle) {
	b.bounds = bounds

	selectorWidth := b.selector.img.Bounds().Dx()
	selectorHeight := b.selector.img.Bounds().Dy()
	selectorHeightHalf := selectorHeight / 2
	selectorX := b.bounds.Min.X + 20
	selectorY := b.bounds.Min.Y + (b.bounds.Dy() / 2) - selectorHeightHalf
	b.selector.Layout(image.Rect(selectorX, selectorY, selectorX+selectorWidth, selectorY+selectorHeight))

	// Text layout
	textX := (bounds.Min.X + bounds.Max.X) / 2
	b.Text.Layout(image.Rect(textX, selectorY+selectorHeightHalf, bounds.Max.X, bounds.Max.Y))
}

func (b *Button) Draw(screen *ebiten.Image) {
	//vector.StrokeRect(screen, float32(posX), float32(posY), float32(width), float32(height), 1, color.White, false)

	if b.IsSelected {
		b.selector.Draw(screen)

		//vector.FillCircle(screen, float32(selectorX), float32(selectorY), config.PRESS_START_2P_SIZE/2, color.White, false)
	}

	// Label of button
	b.Text.Draw(screen)
}
