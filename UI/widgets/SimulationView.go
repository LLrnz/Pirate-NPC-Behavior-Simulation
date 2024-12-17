package widgets

import (
	"image/color"
	"log"

	"gitlab.utc.fr/arsaintm/ia04-projet/UI/utils"
	"gitlab.utc.fr/arsaintm/ia04-projet/UI/views"
	"gitlab.utc.fr/arsaintm/ia04-projet/simulation"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

// Widget used for displaying the simulation : the world map with
// all the objetcs (mostly ships) on it
type SimulationView struct {
	*widget.Container
	background *ebiten.Image
	factory    *views.ViewFactory
	objects    []views.IViewAgent
	events     chan Events
}

func NewSimulationView(events chan Events) *SimulationView {
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{255, 0, 0, 255})),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(utils.GameWidth, utils.GameHeight),
		),
	)

	//Load the background image
	fond, err := utils.NewImageFromPath(utils.MapImage)
	if err != nil {
		log.Fatal(err)
	}

	factory, err := views.NewViewFactory()
	if err != nil {
		log.Fatal(err)
	}

	// Converts all simulation objects to a displayable object
	viewAgents, err := factory.ConvertAllAgentsToView([]simulation.Agent{})
	if err != nil {
		log.Fatal(err)
	}

	return &SimulationView{rootContainer, fond, factory, viewAgents, events}
}

// Refreshes the view
func (s *SimulationView) Update(simu *simulation.Simulation) {

	viewAgents, err1 := s.factory.ConvertAllAgentsToView(simu.Agents())
	if err1 != nil {
		log.Fatal(err1)
	}

	s.objects = viewAgents
}

// Draws the background and all objetcs on screen
func (s *SimulationView) Render(screen *ebiten.Image) {

	//Draw the background at an offset of 0,0
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(s.background, op)

	for _, obj := range s.objects {
		obj.Draw(screen)
	}
}
