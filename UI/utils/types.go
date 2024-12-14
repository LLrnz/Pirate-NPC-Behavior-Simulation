package utils

import (
	"image"

	"gitlab.utc.fr/arsaintm/ia04-projet/simulation"
)

// Represents the position on screen (in pixels)
type Position struct {
	image.Point
}

func NewPosition(x int, y int) *Position {
	return &Position{image.Point{x, y}}
}

// Converts a simulation coordiantes to a point on screen
func ToPos(point simulation.Coordinates) Position {
	return *NewPosition(point.X()*TileSize, point.Y()*TileSize)
}
