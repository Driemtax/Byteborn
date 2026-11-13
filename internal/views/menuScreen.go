package views

import (
	"image"

	"github.com/Driemtax/Byteborn/internal/config"
	"github.com/Driemtax/Byteborn/pkg/input"
	"github.com/Driemtax/Byteborn/pkg/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	PADDING = 20
)

type MenuActions struct {
	OnStart    func()
	OnSettings func()
	OnQuit     func()
}

type MenuScreen struct {
	titleDiv      *ui.Div
	logoDiv       *ui.Div
	ButtonManager *ui.ButtonManager

	actions MenuActions
}

func NewMenuScreen(actions MenuActions) *MenuScreen {
	menuScreen := &MenuScreen{
		actions: actions,
	}

	menuScreen.Init()

	return menuScreen
}

func (menu *MenuScreen) Init() {
	// --- TITLE DIV ---
	titleWidth := (config.WINDOW_WIDTH/3)*2 - 2*PADDING
	titleHeight := config.WINDOW_HEIGHT / 3
	titleX := (config.WINDOW_WIDTH / 3) + PADDING
	titleY := 20
	titleText := ui.NewText("Byteborn", text.AlignCenter, text.AlignCenter)
	titleDiv := ui.NewDiv(true)
	titleDiv.AddChild(titleText)
	titleDiv.Layout(image.Rect(titleX, titleY, titleX+titleWidth, titleY+titleHeight))
	menu.titleDiv = titleDiv
	// -----------------------

	// --- LOGO DIV ---
	logoWidth := (config.WINDOW_WIDTH / 3) - PADDING
	logoHeight := config.WINDOW_HEIGHT / 3
	logoX := 20
	logoY := 20
	logoDiv := ui.NewDiv(false)

	pic := ui.NewPicture(config.TITLE_LOGO_PATH)
	logoDiv.AddChild(pic)
	logoDiv.Layout(image.Rect(logoX, logoY, logoX+logoWidth, logoY+logoHeight))
	menu.logoDiv = logoDiv
	// -------------------------------

	// --- BUTTONS ---
	selectionWidth := config.WINDOW_WIDTH / 3
	selectionHeight := (config.WINDOW_HEIGHT/3)*2 - 3*PADDING
	selectionX := (config.WINDOW_WIDTH / 2) - (selectionWidth / 2)
	selectionY := titleY + titleHeight + PADDING
	bm := ui.NewButtonManager()
	bmRect := image.Rect(selectionX, selectionY, selectionX+selectionWidth, selectionY+selectionHeight)
	bm.AddNewButton("Start", true, menu.actions.OnStart)
	bm.AddNewButton("Settings", false, menu.actions.OnSettings)
	bm.AddNewButton("Quit", false, menu.actions.OnQuit)
	bm.Layout(bmRect)
	menu.ButtonManager = bm
}

func (menu *MenuScreen) Update(input *input.InputState) {
	if input.ARROW_UP {
		menu.ButtonManager.SelectPrevious()
	} else if input.ARROW_DOWN {
		menu.ButtonManager.SelectNext()
	}

	if input.ENTER {
		menu.ButtonManager.ActivateSelected()
	}
}

func (menu *MenuScreen) Draw(screen *ebiten.Image) {
	menu.logoDiv.Draw(screen)
	menu.titleDiv.Draw(screen)
	menu.ButtonManager.Draw(screen)
}
