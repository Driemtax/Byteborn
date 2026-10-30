package views

import (
	"github.com/Driemtax/Byteborn/internal/config"
	"github.com/Driemtax/Byteborn/pkg/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	PADDING = 20
)

type MenuScreen struct {
	root          *ui.Div
	titleDiv      *ui.Div
	logoDiv       *ui.Div
	buttonManager *ui.ButtonManager
}

func NewMenuScreen() *MenuScreen {
	return &MenuScreen{
		root: ui.NewDiv(config.WINDOW_WIDTH, config.WINDOW_HEIGHT, 0, 0, false),
	}
}

func (menu *MenuScreen) Init() {
	// --- TITLE DIV ---
	titleWidth := (config.WINDOW_WIDTH/3)*2 - 2*PADDING
	titleHeight := config.WINDOW_HEIGHT / 3
	titleX := (config.WINDOW_WIDTH / 3) + PADDING
	titleY := 20
	titleText := ui.NewText("Byteborn", text.AlignCenter, text.AlignCenter)
	titleDiv := ui.NewDiv(titleWidth, titleHeight, titleX, titleY, true)
	titleDiv.AddChild(titleText)
	menu.titleDiv = titleDiv
	// -----------------------

	// --- LOGO DIV ---
	logoWidth := (config.WINDOW_WIDTH / 3) - PADDING
	logoHeight := config.WINDOW_HEIGHT / 3
	logoX := 20
	logoY := 20
	logoDiv := ui.NewDiv(logoWidth, logoHeight, logoX, logoY, false)

	pic := ui.NewPicture(config.TITLE_LOGO_PATH)
	logoDiv.AddChild(pic)
	menu.logoDiv = logoDiv
	// -------------------------------

	// --- BUTTONS ---
	selectionWidth := config.WINDOW_WIDTH / 3
	selectionHeight := (config.WINDOW_HEIGHT/3)*2 - 3*PADDING
	selectionX := (config.WINDOW_WIDTH / 2) - (selectionWidth / 2)
	selectionY := titleY + titleHeight + PADDING
	bm := ui.NewButtonManager(selectionWidth, selectionHeight, selectionX, selectionY)
	bm.AddNewButton("Start", true)
	bm.AddNewButton("Settings", false)
	bm.AddNewButton("Quit", false)
	menu.buttonManager = bm
}

func (menu *MenuScreen) Draw(screen *ebiten.Image, posX int, posY int, width int, height int) {
	menu.root.Draw(screen)
	menu.logoDiv.Draw(screen)
	menu.titleDiv.Draw(screen)
	menu.buttonManager.Draw(screen, posX, posY, width, height)
}
