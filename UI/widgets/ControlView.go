package widgets

import (
	"image/color"

	"gitlab.utc.fr/arsaintm/ia04-projet/UI/utils"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

// Widget used for displaying buttons to control the simulation
type ControlView struct {
	*widget.Container
	events chan Events
	btn    *widget.Button
}

func NewControlView(events chan Events) *ControlView {
	// Create UI container
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.RGBA{100, 100, 0, 255})),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(100, 50),
		),
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(1),

			widget.GridLayoutOpts.Stretch([]bool{true}, []bool{true}),
		)),
	)

	buttonImage, _ := utils.LoadButtonImage()

	// load button text font
	face, _ := utils.LoadFont(20)

	// Creates the "Add Agent" button
	var button *widget.Button
	button = widget.NewButton(
		widget.ButtonOpts.WidgetOpts(),
		widget.ButtonOpts.Image(buttonImage),

		widget.ButtonOpts.Text("Add ship", face, &widget.ButtonTextColor{
			Idle:    color.NRGBA{0xdf, 0xf4, 0xff, 0xff},
			Hover:   color.NRGBA{0, 255, 128, 255},
			Pressed: color.NRGBA{255, 0, 0, 255},
		}),
		widget.ButtonOpts.TextProcessBBCode(true),
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:   30,
			Right:  30,
			Top:    5,
			Bottom: 5,
		}),

		//When button presses
		widget.ButtonOpts.PressedHandler(func(args *widget.ButtonPressedEventArgs) {
			button.Text().Inset.Top = 4
			button.Text().Inset.Left = 4
			button.GetWidget().CustomData = true
		}),

		//When button released
		widget.ButtonOpts.ReleasedHandler(func(args *widget.ButtonReleasedEventArgs) {
			button.Text().Inset.Top = 0
			button.Text().Inset.Left = 0
			button.GetWidget().CustomData = false
		}),

		// When button clicked
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			events <- AddAgent
		}))

	rootContainer.AddChild(button)

	return &ControlView{rootContainer, events, button}
}
