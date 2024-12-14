package UI

import (
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"gitlab.utc.fr/arsaintm/ia04-projet/UI/controllers"
	"gitlab.utc.fr/arsaintm/ia04-projet/UI/widgets"
	"gitlab.utc.fr/arsaintm/ia04-projet/simulation"
)

func NewMainWindow(model *simulation.Simulation, refreshRate float64) (*widget.Container, controllers.SimulationController) {
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

	//rootroot.AddChild(rootContainer)
	eventChan := make(chan widgets.Events)
	controlView := widgets.NewControlView(eventChan)
	viewSimu := widgets.NewSimulationView(eventChan)

	controller := controllers.NewSimulationController(model, viewSimu, controlView, eventChan, refreshRate)

	rootroot.AddChild(
		viewSimu,
		widgets.NewLogView(),
		widgets.NewInfoPanelView(),
		controlView,
		//innerContainer2,
	)
	return rootroot, *controller
}
