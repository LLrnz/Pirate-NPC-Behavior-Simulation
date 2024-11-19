package ui

import (
	"fmt"
	"image"
	"simulationPirate/pkg/engine"
	"simulationPirate/pkg/engine/grid"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func Update() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		obj, ok := GetObjectAtCursor()
		if ok && *obj != nil {
			fmt.Print((*obj).ToString())
		} else {
			fmt.Print("this cell is empty")
		}

	}
}

func GetObjectAtCursor() (*engine.Object, bool) {
	x, y := ebiten.CursorPosition()
	cursorPosition := engine.Position{Point: image.Point{x, y}}

	return grid.GetInstance().GetObjectAtPos(cursorPosition)
}
