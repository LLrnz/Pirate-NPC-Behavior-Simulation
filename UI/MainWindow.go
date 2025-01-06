package UI

import (
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"gitlab.utc.fr/arsaintm/ia04-projet/UI/controllers"
	"gitlab.utc.fr/arsaintm/ia04-projet/UI/widgets"
	"gitlab.utc.fr/arsaintm/ia04-projet/simulation"
)

// Instantiate the Window container window and all its children
func NewMainWindow(model *simulation.Simulation, refreshRate float64) (*widget.Container, controllers.SimulationController) {
	rootroot := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(color.NRGBA{0x13, 0x1a, 0x22, 0xff}),
		),
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(2),
			widget.GridLayoutOpts.Stretch([]bool{true, false}, []bool{true}),
		)),
	)

	eventChan := make(chan widgets.Events)

	// the control view
	controlView := widgets.NewControlView(model)
	infoPanelView := widgets.NewInfoPanelView()
	// the simulation view
	viewSimu := widgets.NewSimulationView(model, eventChan, infoPanelView)

	//creates the view controller
	controller := controllers.NewSimulationController(model, viewSimu, controlView, eventChan, refreshRate)

	// Add all children to window in row
	rootroot.AddChild(
		viewSimu,
		widgets.NewLogView(model),
	)
	return rootroot, *controller
}
