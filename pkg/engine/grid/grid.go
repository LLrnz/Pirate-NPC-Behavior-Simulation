package grid

import (
	"fmt"
	"image"
	"simulationPirate/pkg/engine"
	"simulationPirate/pkg/engine/sprite"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

var Right image.Point = image.Point{1, 0}
var Left image.Point = image.Point{-1, 0}
var Up image.Point = image.Point{0, -1}
var Down image.Point = image.Point{0, 1}

type Grid struct {
	Grid   map[image.Point]engine.Object //ObjectMap[image.Point, engine.Object]
	Height int
	Width  int
}

var instance *Grid
var once sync.Once
var Initialized bool = false

// GetInstance returns the singleton instance (thread-safe)
func GetInstance() *Grid {
	firstCall := false
	once.Do(func() {
		// Initialize the instance only once
		instance = &Grid{map[image.Point]engine.Object{}, 0, 0}
		firstCall = true
	})
	if !Initialized && !firstCall {
		fmt.Print("Grid not initialized")
		return nil
	}
	return instance
}

func (g *Grid) Init(height int, width int) {
	g.Height = height
	g.Width = width
	Initialized = true
}

func (g *Grid) GetNeighbors(cell image.Point) []image.Point {
	neighbors := []image.Point{}

	//si n'est pas collé au bord gauche
	if cell.X > 0 {
		neighbors = append(neighbors, cell.Add(Left))

		//sinon si n'est pas collé au bord droit
	} else if cell.X < g.Width {
		neighbors = append(neighbors, cell.Add(Right))
	}

	//si n'est pas collé au bord haut
	if cell.Y > 0 {
		neighbors = append(neighbors, cell.Add(Up))

		//sinon si n'est pas collé au bord bas
	} else if cell.Y < g.Height {
		neighbors = append(neighbors, cell.Add(Down))
	}

	return neighbors
}

func (g *Grid) GetNeighbor(position image.Point, direction sprite.Direction) image.Point {
	cell := ToCoor(position)
	switch direction {
	case sprite.Up:
		if cell.Y > 0 {
			return cell.Add(Up)
		}
	case sprite.Down:
		if cell.Y < g.Height {
			return cell.Add(Down)
		}
	case sprite.Right:
		if cell.X < g.Width {
			return cell.Add(Right)
		}
	case sprite.Left:
		if cell.X > 0 {
			return cell.Add(Left)

		}
	}

	return cell
}

func (g *Grid) Update() {
	for _, obj := range g.Grid {
		obj.Update()
	}
}

func (g *Grid) Draw(screen *ebiten.Image) {
	for _, obj := range g.Grid {
		obj.Draw(screen)
	}
}

func ToPos(point image.Point) image.Point {
	return point.Mul(engine.TileSize)
}

func ToCoor(point image.Point) image.Point {
	return point.Div(engine.TileSize)
}
