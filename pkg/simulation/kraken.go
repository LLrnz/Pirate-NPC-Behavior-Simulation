package simulation

import (
	"fmt"
	"image"
	"simulationPirate/pkg/engine/grid"
	"simulationPirate/pkg/engine/sprite"
)

type Kraken struct {
	sprite.Sprite
}

func NewKraken(position image.Point) *Kraken {

	sprite, err := sprite.NewSprite(position, sprite.Right, 1, "assets/bateau.png")
	if err != nil {
		fmt.Printf("création de Kraken impossible : %s", err)
	}

	return &Kraken{*sprite}
}

func (p *Kraken) Update() {
	next := grid.ToPos(grid.GetInstance().GetNeighbor(p.Position, p.Direction))
	p.Position = next
	fmt.Print(next)
}
