package main

import (
	"image/color"
	"log"
	"simulationPirate/pkg/controllers"
	"simulationPirate/pkg/models"
	"simulationPirate/pkg/views"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

const maxTick int = 2 * 60

type game struct {
	ui         *ebitenui.UI
	ticker     int
	controller controllers.SimulationController
	model      *models.Simulation
}

// Layout implements Game.
func (g *game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *game) Update() error {
	g.ui.Update()

	if g.ticker >= maxTick {
		g.controller.Update(g.model)
		g.ticker = 0

	}
	g.ticker++

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	g.ui.Draw(screen)
}

func main() {

	// construct a new container that serves as the root of the UI hierarchy
	rootroot := widget.NewContainer(
		// the container will use a plain color as its background
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{0x13, 0x1a, 0x22, 0xff})),
		// the container will use an anchor layout to layout its single child widget
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(2),

			widget.GridLayoutOpts.Stretch([]bool{false, true}, []bool{false, true}),
		//widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
		)),
	)

	model := models.Simulation{}

	//rootroot.AddChild(rootContainer)
	eventChan := make(chan views.Events)
	controlView := views.NewControlView(eventChan)
	viewSimu := views.NewSimulationView(eventChan)

	controller := controllers.NewSimulationController(&model, viewSimu, controlView, eventChan)

	rootroot.AddChild(
		viewSimu,
		views.NewLogView(),
		views.NewInfoPanelView(),
		controlView,
		//innerContainer2,
	)

	// construct the UI
	ui := ebitenui.UI{
		Container: rootroot,
	}

	// Ebiten setup
	ebiten.SetWindowSize(868, 432)
	ebiten.SetWindowSizeLimits(968, 532, -1, -1)
	ebiten.SetWindowTitle("Simulation")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	game := game{
		ui:         &ui,
		controller: *controller,
		model:      &model,
	}
	//game.LoadContent()

	// run Ebiten main loop
	err := ebiten.RunGame(&game)
	if err != nil {
		log.Println(err)
	}
}
