package engine

import (
	"image"

	"gitlab.utc.fr/arsaintm/ia04-projet/simulation"
)

type Position struct {
	image.Point
}

func NewPosition(x int, y int) *Position {
	return &Position{image.Point{x, y}}
}

func ToPos(point simulation.Coordinates) Position {
	return *NewPosition(point.X()*TileSize, point.Y()*TileSize)
}
