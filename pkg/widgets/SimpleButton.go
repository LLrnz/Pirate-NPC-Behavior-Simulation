package widgets

import (
	"image/color"

	"testGUI/pkg/utils"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

type SimpleButton struct {
	*widget.Button
	OnClick func()
}

func NewSimpleButton(OnClick func()) *widget.Button {
	buttonImage, _ := LoadButtonImage()

	// load button text font
	face, _ := utils.LoadFont(20)

	// Creating button variable first so that it is usable in callbacks
	var button *widget.Button
	// construct a button
	button = widget.NewButton(
		// set general widget options
		widget.ButtonOpts.WidgetOpts(
			// instruct the container's anchor layout to center the button both horizontally and vertically
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
		),
		// specify the images to use
		widget.ButtonOpts.Image(buttonImage),

		// specify the button's text, the font face, and the color
		//widget.ButtonOpts.Text("Hello, World!", face, &widget.ButtonTextColor{
		widget.ButtonOpts.Text("Hello, [color=FF00FF]World![/color]", face, &widget.ButtonTextColor{
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
			OnClick()
		}),

		// // add a handler that reacts to entering the button with the cursor
		// widget.ButtonOpts.CursorEnteredHandler(func(args *widget.ButtonHoverEventArgs) {
		// 	//println("cursor entered button: entered =", args.Entered, "offsetX =", args.OffsetX, "offsetY =", args.OffsetY)
		// 	//If we moved the Text because we clicked on this button previously, move the text down and right
		// 	if button.GetWidget().CustomData == true {
		// 		button.Text().Inset.Top = 4
		// 		button.Text().Inset.Left = 4
		// 	}
		// }),

		// // add a handler that reacts to moving the cursor on the button
		// widget.ButtonOpts.CursorMovedHandler(func(args *widget.ButtonHoverEventArgs) {
		// 	//println("cursor moved on button: entered =", args.Entered, "offsetX =", args.OffsetX, "offsetY =", args.OffsetY, "diffX =", args.DiffX, "diffY =", args.DiffY)
		// }),

		// // add a handler that reacts to exiting the button with the cursor
		// widget.ButtonOpts.CursorExitedHandler(func(args *widget.ButtonHoverEventArgs) {
		// 	//println("cursor exited button: entered =", args.Entered, "offsetX =", args.OffsetX, "offsetY =", args.OffsetY)
		// 	//Reset the Text inset if the cursor is no longer over the button
		// 	button.Text().Inset.Top = 0
		// 	button.Text().Inset.Left = 0
		// }),
	)
	return button
}

func LoadButtonImage() (*widget.ButtonImage, error) {
	idle := image.NewNineSliceColor(color.NRGBA{R: 170, G: 170, B: 180, A: 255})

	hover := image.NewNineSliceColor(color.NRGBA{R: 130, G: 130, B: 150, A: 255})

	pressed := image.NewNineSliceColor(color.NRGBA{R: 100, G: 100, B: 120, A: 255})

	return &widget.ButtonImage{
		Idle:    idle,
		Hover:   hover,
		Pressed: pressed,
	}, nil
}
