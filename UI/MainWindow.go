package UI

import (
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"gitlab.utc.fr/arsaintm/ia04-projet/UI/controllers"
	"gitlab.utc.fr/arsaintm/ia04-projet/UI/widgets"
	"gitlab.utc.fr/arsaintm/ia04-projet/simulation"
)

// Instantiate the container window and all its children
func NewMainWindow(model *simulation.Simulation, refreshRate float64) (*widget.Container, controllers.SimulationController) {
	// construct a new container that serves as the root of the UI hierarchy
	rootroot := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{0x13, 0x1a, 0x22, 0xff})),
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(2),

			widget.GridLayoutOpts.Stretch([]bool{false, true}, []bool{false, true}),
		)),
	)

	eventChan := make(chan widgets.Events)
	// Create the children
	// the control view
	controlView := widgets.NewControlView(eventChan)
	// the simulation view
	viewSimu := widgets.NewSimulationView(eventChan)

	//creates the view controller
	controller := controllers.NewSimulationController(model, viewSimu, controlView, eventChan, refreshRate)

	// Add all children to window
	rootroot.AddChild(
		viewSimu,
		widgets.NewLogView(),
		widgets.NewInfoPanelView(),
		controlView,
	)
	return rootroot, *controller
}
