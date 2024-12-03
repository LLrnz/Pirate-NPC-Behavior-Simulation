package views

import (
	im "image"
	"image/color"
	"log"

	"testGUI/pkg/engine"
	"testGUI/pkg/engine/sprite"
	"testGUI/pkg/models"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

type SimulationView struct {
	*widget.Container
	fond    *sprite.Sprite
	objects []engine.Object
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

	return &SimulationView{rootContainer, fond, []engine.Object{}, events}
}

func (s *SimulationView) Update(simu *models.Simulation) {
	s.objects = simu.Objects
}

func (s *SimulationView) Render(screen *ebiten.Image) {
	s.fond.Draw(screen)

	for _, obj := range s.objects {
		obj.Draw(screen)
	}
}
