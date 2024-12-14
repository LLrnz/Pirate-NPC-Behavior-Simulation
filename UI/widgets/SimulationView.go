package widgets

import (
	im "image"
	"image/color"
	"log"

	"gitlab.utc.fr/arsaintm/ia04-projet/UI/engine"
	"gitlab.utc.fr/arsaintm/ia04-projet/UI/engine/sprite"
	"gitlab.utc.fr/arsaintm/ia04-projet/UI/views"
	"gitlab.utc.fr/arsaintm/ia04-projet/simulation"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

type SimulationView struct {
	*widget.Container
	fond    *sprite.Sprite
	objects []views.IViewAgent
	events  chan Events
}

func NewSimulationView(events chan Events) *SimulationView {
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{255, 0, 0, 255})),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(768, 432),
		),
	)

	fond, err := sprite.NewSprite(engine.Position{Point: im.Point{0, 0}}, sprite.Right, 4, engine.MapImage)
	if err != nil {
		log.Fatal(err)
	}

	viewAgents, err1 := views.ConvertAllAgentsToView([]simulation.Agent{})
	if err1 != nil {
		log.Fatal(err1)
	}

	return &SimulationView{rootContainer, fond, viewAgents, events}
}

func (s *SimulationView) Update(simu *simulation.Simulation) {

	viewAgents, err1 := views.ConvertAllAgentsToView(simu.Agents())
	if err1 != nil {
		log.Fatal(err1)
	}

	s.objects = viewAgents
}

func (s *SimulationView) Render(screen *ebiten.Image) {
	s.fond.Draw(screen)

	for _, obj := range s.objects {
		obj.Draw(screen)
	}
}
