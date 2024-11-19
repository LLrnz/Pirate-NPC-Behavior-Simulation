package ui

import (
	"bytes"
	"image/color"
	"simulationPirate/pkg/engine"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var font, _ = text.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2P_ttf))

type InfoPanel struct {
	position      engine.Position
	currentObject *engine.Object
	hasObject     bool
	//font          *text.GoTextFaceSource
}

func (i *InfoPanel) ChangeFocus(obj *engine.Object, exists bool) {
	if exists {
		i.currentObject = obj
		i.hasObject = true
	} else {
		i.hasObject = false
	}
}

func (i *InfoPanel) Update() {

}

func (i *InfoPanel) Draw(screen *ebiten.Image) {

	if i.hasObject {

		//background
		opt := &text.DrawOptions{}
		opt.GeoM.Translate(float64(i.position.X), float64(i.position.Y))
		opt.ColorScale.ScaleWithColor(color.White)
		rectangle := ebiten.NewImage(100, 100)
		rectangle.Fill(color.Black)

		//text

		msg := (*i.currentObject).ToString()
		op := &text.DrawOptions{}
		//op.GeoM.Translate(float64(i.position.X), float64(i.position.Y))
		op.ColorScale.ScaleWithColor(color.White)
		text.Draw(rectangle, msg, &text.GoTextFace{
			Source: font,
			Size:   8,
		}, op)

		screen.DrawImage(rectangle, &opt.DrawImageOptions)

	}

}

func (i *InfoPanel) Position() engine.Position {
	return i.position
}

func (i *InfoPanel) SetPosition(p engine.Position) {
	i.position = p
}
