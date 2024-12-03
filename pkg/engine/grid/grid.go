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
	Grid     map[engine.Coordinates]engine.Object //ObjectMap[image.Point, engine.Object]
	Height   int
	Width    int
	toAdd    map[engine.Coordinates]engine.Object
	toRemove map[engine.Coordinates]engine.Object
}

var instance *Grid
var once sync.Once
var Initialized bool = false

// GetInstance returns the singleton instance (thread-safe)
func GetInstance() *Grid {
	firstCall := false
	once.Do(func() {
		// Initialize the instance only once
		instance = &Grid{map[engine.Coordinates]engine.Object{}, 0, 0, map[engine.Coordinates]engine.Object{}, map[engine.Coordinates]engine.Object{}}
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

func (g *Grid) GetNeighbor(position engine.Position, direction sprite.Direction) engine.Coordinates {
	cell := ToCoor(position)
	var point image.Point

	switch direction {
	case sprite.Up:
		if cell.Y > 0 {
			point = cell.Add(Up)
		}
	case sprite.Down:
		if cell.Y < g.Height {
			point = cell.Add(Down)
		}
	case sprite.Right:
		if cell.X < g.Width {
			point = cell.Add(Right)
		}
	case sprite.Left:
		if cell.X > 0 {
			point = cell.Add(Left)

		}
	}

	return engine.Coordinates{Point: point}
}

func (g *Grid) Update() {
	for _, obj := range g.Grid {
		if obj != nil {
			obj.Update()

		}
	}

	g.applyChanges()
}

func (g *Grid) Draw(screen *ebiten.Image) {
	for _, obj := range g.Grid {
		if obj != nil {
			obj.Draw(screen)

		}
	}
}

func ToPos(point engine.Coordinates) engine.Position {
	return engine.Position{Point: point.Mul(engine.TileSize)}
}

func ToCoor(point engine.Position) engine.Coordinates {

	remainderX, remainderY := point.X%engine.TileSize, point.Y%engine.TileSize
	x, y := (point.X-remainderX)/engine.TileSize, (point.Y-remainderY)/engine.TileSize
	//fmt.Printf("clicked on (coor) (%d,%d)/n", x, y)
	return engine.Coordinates{Point: image.Point{x, y}}
}

func (g *Grid) MoveAt(obj engine.Object, new_pos engine.Position) {
	g.RemoveObject(obj.Position(), obj)
	g.SetObjectAtPos(new_pos, obj)
	obj.SetPosition(new_pos)
}

func (g *Grid) GetObjectAt(position engine.Position) (*engine.Object, bool) {
	coor := ToCoor(position)

	object, ok := g.Grid[coor]
	return &object, ok
}

func (g *Grid) MoveNext(obj engine.Object, direction sprite.Direction) {
	g.MoveAt(obj, ToPos(g.GetNeighbor(obj.Position(), direction)))
}

func (g *Grid) SetObjectAtPos(key engine.Position, value engine.Object) {
	g.toAdd[ToCoor(key)] = value
}

func (g *Grid) GetObjectAtPos(key engine.Position) (*engine.Object, bool) {
	object, ok := g.Grid[ToCoor(key)]
	if ok {
		return &object, ok
	}
	return nil, ok
}

func (g *Grid) RemoveObject(key engine.Position, value engine.Object) {
	g.toRemove[ToCoor(key)] = value
}

func (g *Grid) applyChanges() {
	for key := range g.toRemove {
		g.Grid[key] = nil
	}

	for key, val := range g.toAdd {
		g.Grid[key] = val
	}

	g.toAdd = map[engine.Coordinates]engine.Object{}
	g.toRemove = map[engine.Coordinates]engine.Object{}

}
