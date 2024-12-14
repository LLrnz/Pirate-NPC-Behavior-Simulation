package widgets

import (
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

// Widget used to display infos about a selected object (WIP)
type InfoPanelView struct {
	*widget.Container
}

func NewInfoPanelView() *InfoPanelView {
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.RGBA{0, 100, 100, 255})),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(100, 50),
		),
	)

	return &InfoPanelView{rootContainer}
}
