package views

import (
	"image/color"
	"testGUI/pkg/utils"
	"testGUI/pkg/widgets"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

type ControlView struct {
	*widget.Container
	events chan Events
	btn    *widget.Button

	//controller engine.SimulationController
}

func NewControlView(events chan Events) *ControlView {
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.RGBA{100, 100, 0, 255})),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(100, 50),
		),
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(1),

			widget.GridLayoutOpts.Stretch([]bool{true}, []bool{true}),
		//widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
		)),
	)

	buttonImage, _ := widgets.LoadButtonImage()

	// load button text font
	face, _ := utils.LoadFont(20)

	// Creating button variable first so that it is usable in callbacks
	var button *widget.Button
	// construct a button
	button = widget.NewButton(
		// set general widget options
		widget.ButtonOpts.WidgetOpts(
		// instruct the container's anchor layout to center the button both horizontally and vertically

		),
		// specify the images to use
		widget.ButtonOpts.Image(buttonImage),

		widget.ButtonOpts.Text("Add ship", face, &widget.ButtonTextColor{
			Idle:    color.NRGBA{0xdf, 0xf4, 0xff, 0xff},
			Hover:   color.NRGBA{0, 255, 128, 255},
			Pressed: color.NRGBA{255, 0, 0, 255},
		}),
		widget.ButtonOpts.TextProcessBBCode(true),
		// specify that the button's text needs some padding for correct display
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:   30,
			Right:  30,
			Top:    5,
			Bottom: 5,
		}),
		//Move the text down and right on press
		widget.ButtonOpts.PressedHandler(func(args *widget.ButtonPressedEventArgs) {
			button.Text().Inset.Top = 4
			button.Text().Inset.Left = 4
			button.GetWidget().CustomData = true
		}),
		//Move the text back to start on press released
		widget.ButtonOpts.ReleasedHandler(func(args *widget.ButtonReleasedEventArgs) {
			button.Text().Inset.Top = 0
			button.Text().Inset.Left = 0
			button.GetWidget().CustomData = false
		}),

		// add a handler that reacts to clicking the button
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			events <- AddAgent
		}))

	rootContainer.AddChild(button)

	return &ControlView{rootContainer, events, button}
}
