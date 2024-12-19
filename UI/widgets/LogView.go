package widgets

import (
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

// Widget to output the logs (WIP)
type LogView struct {
	*widget.Container
}

func NewLogView() *LogView {
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.RGBA{100, 0, 100, 255})),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(100, 50),
		),
	)

	return &LogView{rootContainer}
}
