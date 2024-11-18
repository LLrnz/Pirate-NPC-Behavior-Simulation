package grid

import(
    "simulationPirate/pkg/engine"
	"image"
)

var Right := image.NewPoint(1, 0)
var Left := image.NewPoint(-1, 0)
var Up := image.NewPoint(0, -1)
var Down := image.NewPoint(0, 1)

type Grid struct{
	map[image.Point]engine.Object
	Height int
	Width int
}

func (g *Grid) GetNeighbors(cell image.Point) []image.Point{
	neighbors := []image.Point{}

	//si n'est pas collé au bord gauche
	if cell.X > 0 {
		neighbors = append(neighbors, cell+Left)

	//sinon si n'est pas collé au bord droit
	} else if cell.X < Width {
		neighbors = append(neighbors, cell+Right)
	}

	//si n'est pas collé au bord haut
	if cell.Y > 0 {
		neighbors = append(neighbors, cell+Up)

	//sinon si n'est pas collé au bord bas
	}else if cell.Y < Height{
		neighbors = append(neighbors, cell+Down)
	}

	return neighbors
}