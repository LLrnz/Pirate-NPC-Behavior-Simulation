package ui

import (
	"image"
	"simulationPirate/pkg/engine"
	"simulationPirate/pkg/engine/grid"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type UserInterface struct {
	infoPanel InfoPanel
}

func NewUserInterface() *UserInterface {
	return &UserInterface{InfoPanel{engine.Position{Point: image.Point{0, engine.GameHeight - engine.InfoPanelHeight}}, nil, false}}
}

func (ui *UserInterface) Update() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		obj, ok := ui.GetObjectAtCursor()

		ui.infoPanel.ChangeFocus(obj, ok)

	}
	ui.infoPanel.Update()
}

func (ui *UserInterface) GetObjectAtCursor() (*engine.Object, bool) {
	x, y := ebiten.CursorPosition()
	cursorPosition := engine.Position{Point: image.Point{x, y}}

	return grid.GetInstance().GetObjectAtPos(cursorPosition)
}

func (ui *UserInterface) Draw(screen *ebiten.Image) {
	ui.infoPanel.Draw(screen)
}

func (ui *UserInterface) Position() engine.Position {
	return ui.infoPanel.Position()
}

func (ui *UserInterface) SetPosition(pos engine.Position) {
	ui.infoPanel.SetPosition(pos)
}
