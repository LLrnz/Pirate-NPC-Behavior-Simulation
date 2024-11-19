package simulation

import (
	"fmt"
	"image"
	"simulationPirate/pkg/engine/grid"
	"simulationPirate/pkg/engine/sprite"
)

type Goelette struct {
	sprite.Sprite
}

func NewGoelette(position image.Point) *Goelette {

	sprite, err := sprite.NewSprite(position, sprite.Right, 1, "assets/bateau.png")
	if err != nil {
		fmt.Printf("création de Goelette impossible : %s", err)
		return nil
	}

	return &Goelette{*sprite}
}

func (p *Goelette) Update() {
	next := grid.ToPos(grid.GetInstance().GetNeighbor(p.Position, p.Direction))
	p.Position = next
	fmt.Print(next)
}
